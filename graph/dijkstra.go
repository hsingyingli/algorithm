package graph

import (
	"math"
)

type Record struct {
	visited bool
	cost    int
	from    int
}

func Dijkstra(matrix [][]int, start int) []Record {
	num := len(matrix[0])
	records := make([]Record, num)
	for i := 0; i < num; i++ {
		records[i] = Record{
			visited: i == start,
			cost:    matrix[start][i],
			from:    start,
		}
	}

	for {
		nextNode, ok := next(records)

		if !ok {
			break
		}
		records[nextNode].visited = true

		for i := 0; i < num; i++ {
			cost := sum(records[nextNode].cost, matrix[nextNode][i])

			if cost < records[i].cost {
				records[i].cost = cost
				records[i].from = nextNode
			}
		}
	}

	return records
}

func next(records []Record) (int, bool) {
	nextNode := -1
	min := math.MaxInt
	for i := 0; i < len(records); i++ {
		if !records[i].visited && records[i].cost < min {
			min = records[i].cost
			nextNode = i
		}
	}

	return nextNode, nextNode != -1
}

func sum(a int, b int) int {
	if a == math.MaxInt || b == math.MaxInt {
		return math.MaxInt
	}

	return a + b
}
