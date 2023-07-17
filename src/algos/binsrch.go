package main

import (
	"fmt"
	"math"
)

func main() {
	var searchNum int64 = 8
	lst := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 30, 40, 50, 1000000}
	resulIndex := searchNumberPosition(searchNum, lst)
	if resulIndex == -1 {
		fmt.Printf("Число %d не найдено", searchNum)
	} else {
		fmt.Printf("Число %d найдено на позиции: %d.", searchNum, resulIndex)
	}
}

func searchNumberPosition(num int64, lst []int64) int64 {
	var minIdx int64 = 0
	var maxidx int64 = int64(len(lst) - 1)
	iteration := 0
	for minIdx <= maxidx {
		fmt.Printf("%d \n", iteration)
		midIdx := int64(math.Floor(float64(minIdx+maxidx) / 2))
		midVal := lst[midIdx]

		if num == midVal {
			return maxidx
		}

		if num < midVal {
			maxidx = midIdx - 1
		} else {
			minIdx = midIdx + 1
		}
	}

	return -1
}
