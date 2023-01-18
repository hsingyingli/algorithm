package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandomInt(min int64, max int64) (number int64, ok error) {
	if max < min {
		ok = fmt.Errorf("max must greater than min")
		return
	}
	number = min + rand.Int63n(max-min)
	return
}

func RandomString(length int) (s string, err error) {
	var sb strings.Builder
	var idx int64
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYL"
	k := len(alphabet)

	for i := 0; i < length; i++ {
		idx, err = RandomInt(0, int64(k)-1)
		if err != nil {
			return
		}
		c := alphabet[idx]
		sb.WriteByte(c)
	}

	s = sb.String()

	return
}

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

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() {
	rand.Seed(time.Now().Unix())
}
