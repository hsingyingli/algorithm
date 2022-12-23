package sorting

func bubbleSort(array []int) []int {
	n := len(array)
	for i := 0; i < n; i++ {
		isSorted := true
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
	return array
}
