package kubecost

import (
	"fmt"
	"strings"

	"github.com/kubecost/cost-model/pkg/costmodel/clusters"
	"github.com/kubecost/cost-model/pkg/prom"
	"github.com/kubecost/cost-model/pkg/util/httputil"
)

// FilterCondition is an enum that represents Allocation-specific filtering
// conditions. It is part of the AllocationFilter DSL "grammar".
type FilterCondition int

const (
	ClusterIDEquals FilterCondition = iota
	NamespaceEquals
	NamespaceNotEquals
	LabelEquals
	AnnotationEquals
	NodeEquals
	ControllerKindEquals
	ControllerEquals
	PodEquals
	ServicesContains
	ContainerEquals
)

// AllocationFilter is a mini-DSL for filtering Allocation data by different
// conditions. By specifying a more strict DSL instead of using arbitrary
// functions we gain the ability to take advantage of storage-level filtering
// performance improvements like indexes in databases. We can create a
// transformation from our DSL to the storage's specific query language.
//
// Queries can take filters with []AllocationFilter. Using this approach lets
// us disambiguate between "filter field = empty string" (with
// AllocationFilter{Value: ""}) and "don't filter field" by not specifying
// an AllocationFilter for that field.
//
// TODO: We must be able to translate the user-facing query language (e.g.
// '?filterNamespace=...') to AllocationFilter(s). This includes supporting
// wildcards. This translation should happen in the HTTP query handler.
type AllocationFilterCondition struct {
	Kind  FilterCondition
	Value string
	// ValueSecondary is for multiple comparisons, like checking combined key-
	// value pairs for labels.
	ValueSecondary string
	// WildcardEnd determines if the filter will match any string that starts
	// with Value. If the filter uses ValueSecondary, the wildcard will only
	// apply to ValueSecondary -- Value will be assumed to be an exact match.
	WildcardEnd bool
}

// AllocationFilterOr is a set of filters that should be ORed over instead of
// the default AND. Having a depth-1 OR is necessary for filters like those for
// clusters, where we want to be comparing both cluster ID and cluster name.
type AllocationFilterOr struct {
	Filters []AllocationFilter
}

type AllocationFilterAnd struct {
	Filters []AllocationFilter
}

type AllocationFilter interface {
	Matches(a *Allocation) bool
	// NotEmpty returns true if the filter contains any number of filter
	// conditions. This is necessary for parts of allocation logic like
	// AggregateBy which have conditional logic if there are any filters present.
	NotEmpty() bool
}

func (filter AllocationFilterCondition) NotEmpty() bool {
	return true
}

func (filter AllocationFilterCondition) Matches(a *Allocation) bool {
	if a == nil {
		return false
	}
	if a.Properties == nil {
		return false
	}

	// TODO Controller condition should allow controllerkind:controllername
	// syntax

	switch filter.Kind {
	case ClusterIDEquals:
		if a.Properties.Cluster == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Cluster, filter.Value) {
				return true
			}
		} else if a.Properties.Cluster == filter.Value {
			return true
		}
	case NamespaceEquals:
		if a.Properties.Namespace == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Namespace, filter.Value) {
				return true
			}
		} else if a.Properties.Namespace == filter.Value {
			return true
		}
	case NamespaceNotEquals:
		// TODO: what is the != behavior for unallocated filter?

		if filter.WildcardEnd {
			if !strings.HasPrefix(a.Properties.Namespace, filter.Value) {
				return true
			}
		} else if a.Properties.Namespace != filter.Value {
			return true
		}
	// Comes from GetAnnotation/LabelFilterFunc in KCM
	case LabelEquals:
		val, ok := a.Properties.Labels[filter.Value]
		if (!ok || val == "") && filter.Value == UnallocatedSuffix {
			return true
		} else if ok {
			if filter.WildcardEnd {
				if strings.HasPrefix(val, filter.ValueSecondary) {
					return true
				}
			} else if val == filter.ValueSecondary {
				return true
			}
		}
	case AnnotationEquals:
		val, ok := a.Properties.Annotations[filter.Value]
		if (!ok || val == "") && filter.Value == UnallocatedSuffix {
			return true
		} else if ok {
			if filter.WildcardEnd {
				if strings.HasPrefix(val, filter.ValueSecondary) {
					return true
				}
			} else if val == filter.ValueSecondary {
				return true
			}
		}
	case NodeEquals:
		if a.Properties.Node == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Node, filter.Value) {
				return true
			}
		} else if a.Properties.Node == filter.Value {
			return true
		}
	case ControllerKindEquals:
		if a.Properties.ControllerKind == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.ControllerKind, filter.Value) {
				return true
			}
		} else if a.Properties.ControllerKind == filter.Value {
			return true
		}
	case ControllerEquals:
		if a.Properties.Controller == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Controller, filter.Value) {
				return true
			}
		} else if a.Properties.Controller == filter.Value {
			return true
		}
	case PodEquals:
		if a.Properties.Pod == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Pod, filter.Value) {
				return true
			}
		} else if a.Properties.Pod == filter.Value {
			return true
		}
	case ServicesContains:
		if len(a.Properties.Services) == 0 && filter.Value == UnallocatedSuffix {
			return true
		}

		for _, service := range a.Properties.Services {
			if filter.WildcardEnd {
				if strings.HasPrefix(service, filter.Value) {
					return true
				}
			} else if service == filter.Value {
				return true
			}
		}
	case ContainerEquals:
		if a.Properties.Container == "" && filter.Value == UnallocatedSuffix {
			return true
		}
		if filter.WildcardEnd {
			if strings.HasPrefix(a.Properties.Container, filter.Value) {
				return true
			}
		} else if a.Properties.Container == filter.Value {
			return true
		}
	default:
		// TODO: log an error here? this should never happen
		return false
	}

	return false
}

func (and AllocationFilterAnd) NotEmpty() bool {
	for _, filter := range and.Filters {
		if filter.NotEmpty() {
			return true
		}
	}
	return false
}

func (or AllocationFilterOr) NotEmpty() bool {
	for _, filter := range or.Filters {
		if filter.NotEmpty() {
			return true
		}
	}
	return false
}

// MatchesFilters is the canonical in-Go function for determing if an Allocation
// matches a set of AllocationFilters.
func (and AllocationFilterAnd) Matches(a *Allocation) bool {
	filters := and.Filters
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		if !filter.Matches(a) {
			return false
		}
	}

	return true
}

func (or AllocationFilterOr) Matches(a *Allocation) bool {
	filters := or.Filters
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		if filter.Matches(a) {
			return true
		}
	}

	return false
}

func standardFilterCondition(kind FilterCondition, filterValue string) AllocationFilterCondition {
	return AllocationFilterCondition{
		Kind:        kind,
		Value:       filterValue,
		WildcardEnd: strings.HasSuffix(filterValue, "*"),
	}
}

func kvFilterCondition(kind FilterCondition, filterValue, filterValueSecondary string) AllocationFilterCondition {
	return AllocationFilterCondition{
		Kind:           kind,
		Value:          filterValue,
		ValueSecondary: filterValueSecondary,

		// We only wildcard on the VALUE of the K/V pair
		WildcardEnd: strings.HasSuffix(filterValueSecondary, "*"),
	}
}

// parseKVFilter is responsible for parsing key-value filters from HTTP queries,
// like those that come from "filterLabels=app:test,team:team1"
//
// The values of the returned map should only be length 2
func parseKVFilter(filterValues []string) (map[string][]string, error) {
	result := map[string][]string{}
	for _, raw := range filterValues {
		if raw != "" {
			split := strings.Split(raw, ":")
			if len(split) != 2 {
				return nil, fmt.Errorf("illegal filter: %s", raw)
			}
			aName := prom.SanitizeLabelName(strings.TrimSpace(split[0]))
			aVal := strings.TrimSpace(split[1])
			if _, ok := result[aName]; !ok {
				result[aName] = []string{}
			}
			result[aName] = append(result[aName], aVal)
		}
	}

	return result, nil
}

// FilterFromQuery extracts filters from HTTP query parameters and parses them
// into an AllocationFilter.
func FilterFromQuery(qp httputil.QueryParams, clusterMap clusters.ClusterMap) (AllocationFilter, error) {
	// Query params are a set of AND conditions.
	filter := AllocationFilterAnd{
		Filters: []AllocationFilter{},
	}

	// generate a filter func for each cluster filter, and OR the results
	filterClusters := qp.GetList("filterClusters", ",")
	if len(filterClusters) > 0 {
		clusterFilter := AllocationFilterOr{}
		for _, filter := range filterClusters {
			clusterFilter.Filters = append(clusterFilter.Filters, standardFilterCondition(ClusterIDEquals, filter))

			// TODO: handle the alt. cluster name case, see GetClusterFilterFunc from KCM
		}
		filter.Filters = append(filter.Filters, clusterFilter)
	}

	// generate a filter func for each node filter, and OR the results
	filterNodes := qp.GetList("filterNodes", ",")
	if len(filterNodes) > 0 {
		nodeFilter := AllocationFilterOr{}
		for _, filter := range filterNodes {
			nodeFilter.Filters = append(nodeFilter.Filters, standardFilterCondition(NodeEquals, filter))
		}
		filter.Filters = append(filter.Filters, nodeFilter)
	}

	// generate a filter func for each namespace filter, and OR the results
	filterNamespaces := qp.GetList("filterNamespaces", ",")
	if len(filterNamespaces) > 0 {
		namespaceFilter := AllocationFilterOr{}
		for _, filter := range filterNamespaces {
			namespaceFilter.Filters = append(namespaceFilter.Filters, standardFilterCondition(NamespaceEquals, filter))
		}
		filter.Filters = append(filter.Filters, namespaceFilter)
	}

	// generate a filter func for each controllerKind filter, and OR the results
	filterControllerKinds := qp.GetList("filterControllerKinds", ",")
	if len(filterControllerKinds) > 0 {
		ckFilter := AllocationFilterOr{}
		for _, filter := range filterControllerKinds {
			ckFilter.Filters = append(ckFilter.Filters, standardFilterCondition(ControllerKindEquals, filter))
		}
		filter.Filters = append(filter.Filters, ckFilter)
	}

	// generate a filter func for each controller filter, and OR the results
	filterControllers := qp.GetList("filterControllers", ",")
	if len(filterControllers) > 0 {
		cFilter := AllocationFilterOr{}
		for _, filter := range filterControllers {
			cFilter.Filters = append(cFilter.Filters, standardFilterCondition(ControllerEquals, filter))
		}
		filter.Filters = append(filter.Filters, cFilter)
	}

	// generate a filter func for each pod filter, and OR the results
	filterPods := qp.GetList("filterPods", ",")
	if len(filterPods) > 0 {
		podFilter := AllocationFilterOr{}
		for _, filter := range filterPods {
			podFilter.Filters = append(podFilter.Filters, standardFilterCondition(PodEquals, filter))
		}
		filter.Filters = append(filter.Filters, podFilter)
	}

	// generate a filter func for each annotation filter, and OR the results
	filterAnnotations := map[string][]string{}
	filterAnnotations, err := parseKVFilter(qp.GetList("filterAnnotations", ","))
	if err != nil {
		return nil, fmt.Errorf("unable to parse annotation filters: %s", err)
	}
	if len(filterAnnotations) > 0 {
		annotationFilter := AllocationFilterOr{}
		for _, filter := range filterAnnotations {
			// Enforce to avoid panic, just in case
			// TODO: log?
			if len(filter) != 2 {
				continue
			}
			annotationFilter.Filters = append(annotationFilter.Filters, kvFilterCondition(AnnotationEquals, filter[0], filter[1]))
		}
		filter.Filters = append(filter.Filters, annotationFilter)
	}

	// generate a filter func for each label filter, and OR the results
	filterLabels, err := parseKVFilter(qp.GetList("filterLabels", ","))
	if err != nil {
		return nil, fmt.Errorf("unable to parse label filters: %s", err)
	}
	if len(filterLabels) > 0 {
		labelFilter := AllocationFilterOr{}
		for _, filter := range filterLabels {
			// Enforce to avoid panic, just in case
			// TODO: log?
			if len(filter) != 2 {
				continue
			}
			labelFilter.Filters = append(labelFilter.Filters, kvFilterCondition(LabelEquals, filter[0], filter[1]))
		}
		filter.Filters = append(filter.Filters, labelFilter)
	}

	// generate a filter func for each service filter, and OR the results
	filterServices := qp.GetList("filterServices", ",")
	if len(filterServices) > 0 {
		servicesFilter := AllocationFilterOr{}
		for _, filter := range filterServices {
			servicesFilter.Filters = append(servicesFilter.Filters, standardFilterCondition(ServicesContains, filter))
		}
		filter.Filters = append(filter.Filters, servicesFilter)
	}

	// generate a filter func for each container filter, and OR the results
	filterContainers := qp.GetList("filterContainers", ",")
	if len(filterContainers) > 0 {
		cFilter := AllocationFilterOr{}
		for _, filter := range filterContainers {
			cFilter.Filters = append(cFilter.Filters, standardFilterCondition(ContainerEquals, filter))
		}
		filter.Filters = append(filter.Filters, cFilter)
	}

	return filter, nil
}
