package bst

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
            17
      13           19
  12     16      18
9       15

         9 12 13 15 16 17 18 19
height   0  1  2  0  1  3  0  1
depth    3  2  1  3  2  0  2  1
stb      9  9  9  15 15 9  18 18
sta      9  12 16 15 16 19 18 19
suc      12 13 15 16 17 18 19 -1
pre      -1  9 12 13 15 16 17 18
*/

func createTestBST(t *testing.T) ([]*Node, *Node) {
	values := []int{13, 19, 12, 16, 18, 9, 15}
	root := &Node{value: 17}
	for i := 0; i < len(values); i++ {
		BuildBST(root, values[i])
	}

	nodes := root.traversal()
	actual := make([]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		actual[i] = nodes[i].value
	}
	expect := []int{9, 12, 13, 15, 16, 17, 18, 19}

	require.Equal(t, expect, actual)
	return nodes, root
}

func TestBuildBST(t *testing.T) {
	createTestBST(t)
}
func TestHeight(t *testing.T) {
	nodes, _ := createTestBST(t)
	ans := []int{0, 1, 2, 0, 1, 3, 0, 1}
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].height()
		require.Equal(t, ans[i], actual)
	}
}

func TestDepth(t *testing.T) {
	nodes, _ := createTestBST(t)
	ans := []int{3, 2, 1, 3, 2, 0, 2, 1}
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].depth()
		require.Equal(t, ans[i], actual)
	}
}

func TestSubtreeFirst(t *testing.T) {
	nodes, _ := createTestBST(t)
	ans := []int{9, 9, 9, 15, 15, 9, 18, 18}
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].subtreeFirst()
		require.Equal(t, ans[i], actual.value)
	}
}

func TestSubtreeAfter(t *testing.T) {

	nodes, _ := createTestBST(t)
	ans := []int{9, 12, 16, 15, 16, 19, 18, 19}
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].subtreeAfter()
		require.Equal(t, ans[i], actual.value)
	}
}

func TestSuccessor(t *testing.T) {
	ans := []int{12, 13, 15, 16, 17, 18, 19, -1}
	nodes, _ := createTestBST(t)
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].successor()
		if ans[i] == -1 {
			require.Nil(t, actual)
		} else {
			require.Equal(t, ans[i], actual.value)
		}
	}
}

func TestPredecessor(t *testing.T) {
	ans := []int{-1, 9, 12, 13, 15, 16, 17, 18}
	nodes, _ := createTestBST(t)
	for i := 0; i < len(nodes); i++ {
		actual := nodes[i].predecessor()
		if ans[i] == -1 {
			require.Nil(t, actual)
		} else {
			require.Equal(t, ans[i], actual.value)
		}
	}
}

func TestSubtreeInsertAfter(t *testing.T) {
	nodes, _ := createTestBST(t)
	for i := 0; i < len(nodes); i++ {
		nodes[i].subtreeInsertAfter(&Node{value: i})
		next := nodes[i].successor()
		require.Equal(t, i, next.value)
	}
}

func TestSubtreeInsertABefore(t *testing.T) {
	nodes, _ := createTestBST(t)
	for i := 0; i < len(nodes); i++ {
		nodes[i].subtreeInsertBefore(&Node{value: i})
		prev := nodes[i].predecessor()
		require.Equal(t, i, prev.value)
	}
}

func TestSubtreeDelete(t *testing.T) {
	nodes, root := createTestBST(t)
	origin := []int{9, 12, 13, 15, 16, 17, 18, 19}

	for i := 0; i < len(nodes)-1; i++ {
		currentNodes := root.traversal()
		currentNodes[0].subtreeDelete()
		origin = origin[1:]
		currentNodes = root.traversal()
		require.Equal(t, len(origin), len(currentNodes))
		for i := 0; i < len(origin); i++ {
			require.Equal(t, origin[i], currentNodes[i].value)
		}
	}
}
