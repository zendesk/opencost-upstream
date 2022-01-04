package kubecost

import (
	"strings"
)

// FilterCondition is an enum that represents Allocation-specific filtering
// conditions. It is part of the AllocationFilter DSL "grammar".
type FilterCondition int

const (
	ClusterIDEquals FilterCondition = iota
	NamespaceEquals
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
type AllocationFilter struct {
	Kind  FilterCondition
	Value string
	// WildcardEnd determines if the filter will match any string that starts
	// with Value
	WildcardEnd bool
}

// MatchesFilters is the canonical in-Go function for determing if an Allocation
// matches a set of AllocationFilters.
func MatchesFilters(a *Allocation, filters []AllocationFilter) bool {
	if len(filters) == 0 {
		return true
	}
	if a == nil {
		return true
	}
	if a.Properties == nil {
		return false
	}

	for _, filter := range filters {
		switch filter.Kind {
		case ClusterIDEquals:
			if filter.WildcardEnd {
				if !strings.HasPrefix(a.Properties.Cluster, filter.Value) {
					return false
				}
			} else if a.Properties.Cluster != filter.Value {
				return false
			}
		case NamespaceEquals:
			if filter.WildcardEnd {
				if !strings.HasPrefix(a.Properties.Namespace, filter.Value) {
					return false
				}
			} else if a.Properties.Namespace != filter.Value {
				return false
			}
		default:
			// TODO: log an error here? this should never happen
			return false
		}
	}
	return true
}
