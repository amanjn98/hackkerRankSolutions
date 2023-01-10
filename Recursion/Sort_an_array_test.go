package Recursion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{1, 5, 4, 3, 2}
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, Sort(arr))
}
