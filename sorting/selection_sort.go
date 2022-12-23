package sorting

func prefixMax(array []int, i int) int {
	if i > 0 {
		j := prefixMax(array, i-1)
		if array[j] > array[i] {
			return j
		}
	}
	return i
}

func sort(array []int, i int) {
	if i > 0 {
		j := prefixMax(array, i)
		array[i], array[j] = array[j], array[i]
		sort(array[:], i-1)
	}
}

func selectionSort(array []int) []int {
	sort(array[:], len(array)-1)
	return array
}
