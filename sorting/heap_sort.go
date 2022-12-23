package sorting

func heapify(array []int, n int, i int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left < n && array[left] > array[max] {
		max = left
	}
	if right < n && array[right] > array[max] {
		max = right
	}

	if max != i {
		array[max], array[i] = array[i], array[max]
		heapify(array[:], n, max)
	}
}

// parent: ( i - 1 ) / 2
// left: i * 2 + 1
// right: i * 2 + 2
func buildHeap(array []int, n int) {
	lastNode := n - 1
	parent := (lastNode - 1) / 2
	for i := parent; i >= 0; i-- {
		heapify(array[:], n, i)
	}
}

func heapSort(array []int) []int {
	n := len(array)
	buildHeap(array[:], n)

	for i := n - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapify(array[:], i, 0)
	}

	return array
}
