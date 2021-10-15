package subtract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		name     string
		x, y     float64
		expected float64
	}{
		{"zeros", 0, 0, 0},
		{"+ nums // + res", 5.4, 2.1, 3.3000000000000003},
		{"- nums // + res", -10.05, -7, -3.0500000000000007},
		{"+/- nums // + res", 1, -12, 13},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := subtract(tc.x, tc.y); tc.expected != actual {
				t.Errorf("subtract [%f - %f] failed. wanted: %f got: %f\n", tc.x, tc.y, tc.expected, actual)
			}

			assert.Equal(t, tc.expected, subtract(tc.x, tc.y))
		})
	}
}
