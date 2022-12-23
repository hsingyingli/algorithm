package sorting

func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		j := i - 1
		target := array[i]
		for j >= 0 && target < array[j] {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = target
	}
	return array
}
