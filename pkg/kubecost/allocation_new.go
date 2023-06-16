package kubecost

import (
	"fmt"
	"math"

	filter "github.com/opencost/opencost/pkg/filter21"
	"github.com/opencost/opencost/pkg/filter21/ast"
	"github.com/opencost/opencost/pkg/filter21/matcher"
	"github.com/opencost/opencost/pkg/log"
)

type ResourceValues struct {
	CPU float64
	GPU float64
	RAM float64
}

func (a *Allocation) GetAssetResourcePercentages(totalsSet *AssetTotalsSet, byCluster bool) ResourceValues {
	cpuPct, gpuPct, ramPct := 0.0, 0.0, 0.0

	var ats map[string]*AssetTotals
	var key string
	if byCluster {
		ats = totalsSet.Cluster
		key = a.Properties.Cluster
	} else {
		ats = totalsSet.Node
		key = a.Properties.Node
	}

	if at, ok := ats[key]; ok {
		cpuTotal := at.TotalCPUCost()
		if cpuTotal > 0.0 && !math.IsNaN(cpuTotal) {
			cpuPct = a.CPUTotalCost() / at.TotalCPUCost()
		}

		gpuTotal := at.TotalGPUCost()
		if gpuTotal > 0.0 && !math.IsNaN(gpuTotal) {
			gpuPct = a.GPUTotalCost() / at.TotalGPUCost()
		}

		ramTotal := at.TotalRAMCost()
		if ramTotal > 0.0 && !math.IsNaN(ramTotal) {
			ramPct = a.RAMTotalCost() / at.TotalRAMCost()
		}
	} else {
		log.Debugf("GetAssetResourcePercentages: failed to find asset totals by key for: %s", key)
	}

	return ResourceValues{cpuPct, gpuPct, ramPct}
}

func (as *AllocationSet) aggregate(properties []string, labelConfig *LabelConfig) *AllocationSet {
	results := NewAllocationSet(*as.Window.Start(), *as.Window.End())

	for _, a := range as.Allocations {
		key := a.generateKey(properties, labelConfig)
		a.Name = key
		results.Insert(a)
	}

	return results
}

func (as *AllocationSet) Aggregate(properties []string, labelConfig *LabelConfig) *AllocationSet {
	result := as.Clone()
	return result.aggregate(properties, labelConfig)
}

// TODO do we need this? or can we just put it on the Allocation, itself?
func (as *AllocationSet) ComputeSharingPercentages(sharedAllocationCost float64, ats *AllocationTotalsSet) map[string]float64 {
	result := make(map[string]float64, len(as.Allocations))

	for key, alloc := range as.Allocations {
		// Each allocation should receive, as its percentage, the ratio of its
		// own basic total cost to the total unshared cost of all allocations.
		// So we must remove the allocation's shared and external costs, as
		// well as the shared cost from the totals.
		result[key] = (alloc.TotalCost() - alloc.SharedCost - alloc.ExternalCost) / (ats.TotalCostByCluster() - sharedAllocationCost)
	}

	return result
}

func (as *AllocationSet) distribute(toDistribute *AllocationSet, totalsSet *AssetTotalsSet, byCluster bool) (*AllocationSet, bool) {
	remaining := NewAllocationSet(*as.Window.start, *as.Window.end)

	// TODO

	return remaining, true
}

func (as *AllocationSet) Distribute(toDistribute *AllocationSet, totalsSet *AssetTotalsSet, byCluster bool) (*AllocationSet, *AllocationSet, bool) {
	result := as.Clone()
	remaining, ok := result.distribute(toDistribute, totalsSet, byCluster)
	return result, remaining, ok
}

func (as *AllocationSet) filter(filters filter.Filter, labelConfig *LabelConfig) error {
	var am AllocationMatcher
	if filters == nil {
		am = &matcher.AllPass[*Allocation]{}
	} else {
		compiler := NewAllocationMatchCompiler(labelConfig)
		var err error
		am, err = compiler.Compile(filters)
		if err != nil {
			return fmt.Errorf("compiling filter '%s': %w", ast.ToPreOrderShortString(filters), err)
		}
	}
	if am == nil {
		return fmt.Errorf("unexpected nil filter")
	}

	for key, a := range as.Allocations {
		if !am.Matches(a) {
			delete(as.Allocations, key)
		}
	}

	return nil
}

func (as *AllocationSet) Filter(filters filter.Filter, labelConfig *LabelConfig) (*AllocationSet, error) {
	result := as.Clone()
	err := result.filter(filters, labelConfig)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (as *AllocationSet) scale(coeffs map[string]float64) {
	// TODO
}

func (as *AllocationSet) Scale(coeffs map[string]float64) *AllocationSet {
	result := as.Clone()
	result.scale(coeffs)
	return result
}

func (as *AllocationSet) share(toShare float64, allocTotalsSet *AllocationTotalsSet) {
	// TODO
}

func (as *AllocationSet) Share(toShare float64, allocTotalsSet *AllocationTotalsSet) *AllocationSet {
	result := as.Clone()
	result.share(toShare, allocTotalsSet)
	return result
}

// TODO deal with reconciliation ugh
func (as *AllocationSet) transform(
	aggregationProperties []string,
	idleAllocs *AllocationSet,
	distributeIdle bool,
	distributeByCluster bool,
	shareAllocs *AllocationSet,
	shareCost float64,
	allocTotalsSet *AllocationTotalsSet,
	assetTotalsSet *AssetTotalsSet,
	labelConfig *LabelConfig,
) error {
	undistributed := []*Allocation{}

	// (1) Handle idle
	if distributeIdle {
		remaining, ok := as.distribute(idleAllocs, assetTotalsSet, distributeByCluster)
		if !ok && remaining.Length() > 0 {
			for _, rem := range remaining.Allocations {
				undistributed = append(undistributed, rem)
			}
		}

		remaining, ok = shareAllocs.distribute(idleAllocs, assetTotalsSet, distributeByCluster)
		if !ok && remaining.Length() > 0 {
			for _, rem := range remaining.Allocations {
				undistributed = append(undistributed, rem)
			}
		}
	} else {
		// TODO oi... this might suck
		idleAllocs.scale(nil)

		// Merge
		// TODO
	}

	// (2) Aggregate
	as.aggregate(aggregationProperties, labelConfig)

	// (3) Share costs, both from allocations and not
	for _, share := range shareAllocs.Allocations {
		shareCost += share.TotalCost()
	}
	as.share(shareCost, allocTotalsSet)

	// (4) Insert idle
	if idleAllocs.Length() > 0 {
		for _, idle := range idleAllocs.Allocations {
			// TODO is it this simple? what about naming?
			as.Insert(idle)
		}
	}
	if len(undistributed) > 0 {
		for _, undist := range undistributed {
			// TODO is it this simple? what about naming?
			as.Insert(undist)
		}
	}

	return nil
}

func (as *AllocationSet) Transform(
	aggregationProperties []string,
	idleAllocs *AllocationSet,
	distributeIdle bool,
	distributeByCluster bool,
	shareAllocs *AllocationSet,
	shareCost float64,
	allocTotalsSet *AllocationTotalsSet,
	assetTotalsSet *AssetTotalsSet,
	labelConfig *LabelConfig,
) (*AllocationSet, error) {
	result := as.Clone()
	err := result.transform(
		aggregationProperties,
		idleAllocs,
		distributeIdle,
		distributeByCluster,
		shareAllocs,
		shareCost,
		allocTotalsSet,
		assetTotalsSet,
		labelConfig)
	return result, err
}
