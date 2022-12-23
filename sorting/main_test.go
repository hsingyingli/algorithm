package sorting

import (
	"github/hsingyingli/data-structure-and-algorithm/util"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func testSortingAlgorithm(t *testing.T, fn func(array []int) []int) {
	for i := 0; i < 10; i++ {
		array := util.CreateRandomArray(100)
		sorted := fn(array)
		result := util.IsIncreasing(sorted)
		require.Equal(t, true, result)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
