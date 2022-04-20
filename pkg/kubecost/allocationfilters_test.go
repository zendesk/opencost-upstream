package kubecost

import (
	"testing"
)

func Test_ConditionMatches(t *testing.T) {
	cases := []struct {
		name   string
		a      *Allocation
		filter AllocationFilterCondition

		expected bool
	}{
		{
			name: "NamespaceNotEquals false",
			a: &Allocation{
				Properties: &AllocationProperties{
					Namespace: "kube-system",
				},
			},
			filter: AllocationFilterCondition{
				Kind:  NamespaceNotEquals,
				Value: "kube-system",
			},

			expected: false,
		},
		{
			name: "LabelEquals true",
			a: &Allocation{
				Properties: &AllocationProperties{
					Labels: map[string]string{
						"k1": "v1",
					},
				},
			},
			filter: AllocationFilterCondition{
				Kind:           LabelEquals,
				Value:          "k1",
				ValueSecondary: "v1",
			},

			expected: true,
		},
	}

	for _, c := range cases {
		result := c.filter.Matches(c.a)

		if result != c.expected {
			t.Errorf("%s: expected %t, got %t", c.name, c.expected, result)
		}
	}
}

func Test_ParseExamples(t *testing.T) {
	t.Skip()
	// TODO: These are examples based on an initial proposal of the grammer.
	// Once a design has been decided, these should be updated and then used for
	// tests of parsing

	// Filter: 'label:app=cost-analyzer'
	f1 := AllocationFilterCondition{
		Kind:           LabelEquals,
		Value:          "app",
		ValueSecondary: "cost-analyzer",
	}

	// Filter: 'NOT namespace:kubecost'
	f2 := AllocationFilterCondition{
		Kind:  NamespaceEquals,
		Value: "kubecost",
	}

	// Filter: 'namespace:kube-*'
	f3 := AllocationFilterCondition{
		Kind:        NamespaceEquals,
		Value:       "kube-",
		WildcardEnd: true,
	}

	// Filter: 'annotation:automation.ci.id=abc123'
	f4 := AllocationFilterCondition{
		Kind:           AnnotationEquals,
		Value:          "automation.ci.cd",
		ValueSecondary: "abc123",
	}

	// Filter: 'namespace:kubecost AND label:app=cost-analyzer'
	f5 := AllocationFilterAnd{
		Filters: []AllocationFilter{
			AllocationFilterCondition{
				Kind:  NamespaceEquals,
				Value: "kubecost",
			},
			AllocationFilterCondition{
				Kind:           LabelEquals,
				Value:          "app",
				ValueSecondary: "cost-analyzer",
			},
		},
	}

	// Filter: '(namespace:kubecost OR namespace:kubecost-sec) AND label:app=cost-analyzer'
	f6 := AllocationFilterAnd{[]AllocationFilter{
		AllocationFilterOr{[]AllocationFilter{
			AllocationFilterCondition{
				Kind:  NamespaceEquals,
				Value: "kubecost",
			},
			AllocationFilterCondition{
				Kind:  NamespaceEquals,
				Value: "kubecost-sec",
			},
		}},
		AllocationFilterCondition{
			Kind:           LabelEquals,
			Value:          "app",
			ValueSecondary: "cost-analyzer",
		}},
	}

	// Filter: 'namespace:argo-wf AND pod:"workflowABC UID1"'
	f7 := AllocationFilterAnd{
		Filters: []AllocationFilter{
			AllocationFilterCondition{
				Kind:  NamespaceEquals,
				Value: "argo-wf",
			},
			AllocationFilterCondition{
				Kind:  PodEquals,
				Value: "workflowABC UID1",
			},
		},
	}

	// To avoid compiler errors
	t.Log(f1, f2, f3, f4, f5, f6, f7)
}
