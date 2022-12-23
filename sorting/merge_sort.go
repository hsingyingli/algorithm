package sorting

import "math"

func merge(left []int, right []int) []int {
	l := len(left) + len(right)
	arr := make([]int, l)
	arr1 := make([]int, len(left))
	arr2 := make([]int, len(right))
	arr1 = append(arr1, math.MaxInt)
	arr2 = append(arr2, math.MaxInt)

	i1, i2 := 0, 0

	for i := 0; i < l; i++ {
		if arr1[i1] < arr2[i2] {
			arr[i] = arr1[i1]
			i1++
		} else {
			arr[i] = arr2[i2]
			i2++
		}
	}

	return arr
}

func divide(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := len(array) / 2
	leftArray := array[:middle]
	rightArray := array[middle:]
	return merge(divide(leftArray), divide(rightArray))
}

func mergeSort(array []int) []int {
	return divide(array)
}
