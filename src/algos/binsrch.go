package main

import (
	"fmt"
	"math"
)

func main() {
	var searchNum int64
	fmt.Print("Введи целое число (поиск среди массива от 0 до 20):\n")
	fmt.Scanf("%d\n", &searchNum)

	lst := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
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
	for minIdx <= maxidx {
		midIdx := int64(math.Floor(float64(minIdx+maxidx) / 2))
		midVal := lst[midIdx]

		if num == midVal {
			return midIdx
		}

		if num < midVal {
			maxidx = midIdx - 1
		} else {
			minIdx = midIdx + 1
		}
	}

	return -1
}
