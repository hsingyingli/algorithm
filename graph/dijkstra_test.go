package graph

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

const maxi = math.MaxInt

func createMatrix() [][]int {
	matrix := [][]int{}

	matrix = append(matrix, []int{0, 1, 12, maxi, maxi, maxi})
	matrix = append(matrix, []int{maxi, 0, 9, 3, maxi, maxi})
	matrix = append(matrix, []int{maxi, maxi, 0, maxi, 5, maxi})
	matrix = append(matrix, []int{maxi, maxi, 4, 0, 13, 15})
	matrix = append(matrix, []int{maxi, maxi, maxi, maxi, 0, 4})
	matrix = append(matrix, []int{maxi, maxi, maxi, maxi, maxi, 0})

	return matrix
}

func TestDijkstra(t *testing.T) {
	matrix := createMatrix()
	records := Dijkstra(matrix, 0)
	ans := []int{0, 1, 8, 4, 13, 17}
	for i := 0; i < len(records); i++ {
		require.Equal(t, ans[i], records[i].cost)
	}
}
