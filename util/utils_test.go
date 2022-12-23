package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsIncreasing(t *testing.T) {
	var result bool

	array := []int{1, 2, 3, 4, 5}
	result = IsIncreasing(array)
	require.Equal(t, true, result)

	array = []int{1, 2, 3, 5, 4}
	result = IsIncreasing(array)
	require.Equal(t, false, result)
}
