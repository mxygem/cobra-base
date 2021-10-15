package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		x, y     float64
		expected float64
	}{
		{"zeros", 0, 0, 0},
		{"+ nums // + res", 5.4, 2.1, 7.5},
		{"- nums // + res", -10.05, -7, -17.05},
		{"+/- nums // + res", 1, -12, -11},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if actual := add(tc.x, tc.y); tc.expected != actual {
				t.Errorf("add [%f - %f] failed. wanted: %f got: %f\n", tc.x, tc.y, tc.expected, actual)
			}

			assert.Equal(t, tc.expected, add(tc.x, tc.y))
		})
	}
}
