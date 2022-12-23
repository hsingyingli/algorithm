package util

import (
	"math/rand"
	"time"
)

func IsIncreasing(array []int) bool {
	if len(array) <= 1 {
		return true
	}

	for i := 1; i < len(array); i++ {
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}

func CreateRandomArray(n int) []int {
	return rand.Perm(n)
}

func init() {
	rand.Seed(time.Now().Unix())
}
