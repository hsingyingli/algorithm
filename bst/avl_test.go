package bst

import (
	"github/hsingyingli/data-structure-and-algorithm/util"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildAVL(t *testing.T) {
	for round := 0; round < 100; round++ {

		sequence := util.CreateRandomArray(20)
		log.Println(sequence)
		root := &Node{
			value: sequence[0],
			size:  0,
		}

		for i := 1; i < len(sequence); i++ {
			root = BuildAVL(root, sequence[i])
			skew := root.skew()
			require.Greater(t, 2, skew)
			require.Less(t, -2, skew)
		}
	}

}
