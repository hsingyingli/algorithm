package trie

import (
	"github/hsingyingli/data-structure-and-algorithm/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertTrie(t *testing.T) {
	num := 20

	list := make([]string, num)

	root := &Node{
		dict: make(map[byte]*Node),
		end:  false,
	}

	for i := 0; i < num; i++ {
		word, _ := util.RandomString(10)
		list[i] = word
		root.Insert(word)
	}

	for i := 0; i < num; i++ {
		isExist := root.Find(list[i])
		require.True(t, isExist)
	}

	for i := 0; i < num; i++ {
		word, _ := util.RandomString(10)
		isExist := root.Find(word)
		require.False(t, isExist)
	}
}
