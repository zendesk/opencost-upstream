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
