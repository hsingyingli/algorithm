package sorting

func partition(array []int) int {
	base := array[0]
	pivot := -1
	for i := 0; i < len(array); i++ {
		if array[i] <= base {
			pivot += 1
			array[pivot], array[i] = array[i], array[pivot]
		}
	}

	array[pivot], array[0] = array[0], array[pivot]

	return pivot + 1
}

func sorting(array []int) {
	if len(array) < 2 {
		return
	}

	pivot := partition(array)
	sorting(array[:pivot])
	sorting(array[pivot:])
}

func quickSort(array []int) []int {
	sorting(array[:])
	return array
}
