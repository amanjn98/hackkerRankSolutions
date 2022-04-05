package Heap
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKClosest(t *testing.T) {
	testCases := map[string]struct {
		input    	[][]int
		k			int
		expected 	[][]int
	}{
		"given empty array return empty array": {
			input: [][]int{},
			k:	1,
			expected: [][]int{},
		},
		"given single point array return single point": {
			input: 		[][]int{{1,1}},
			k:			1,
			expected: 	[][]int{{1,1}},
		},
		"given array of points return k closest points": {
			input: 		[][]int{{3,1},{-2,2},{-4,0},{0,1},{-2,-1},{5,-1}},
			k: 			3,
			expected: 	[][]int{{0,1},{-2,-1},{-2,2}},
		},
	}

	for name, v := range testCases {
		tc := v

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := KClosest(tc.input, tc.k)

			assert.Equal(t, actual, tc.expected)
		})
	}
}
