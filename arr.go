package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	items := [5]int64{1, 2, 3, 4, 5}
	var itemsSum int64
	var average float64

	for _, val := range items {
		fmt.Printf("Значение: %d \n", val)
		itemsSum += val
	}
	l := len(items)

	if l > 0 {
		average = float64(itemsSum) / float64(len(items))
	}

	fmt.Printf("Сумма: %d \n", itemsSum)
	fmt.Printf("Среднее: %.2f \n", average)
	fmt.Print("-------------------- \n")
	readFile("arr_digits.txt")
}

func readFile(pth string) {
	file, fOpErr := os.Open(pth)
	if fOpErr != nil {
		log.Fatal(fOpErr)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		trmChars := "\n\r\n\t"
		row := strings.Trim(scanner.Text(), trmChars)
		if row != "" {
			fmt.Printf("Значение из файла: %s\n", row)
		}
	}

	fcErr := file.Close()
	if fcErr != nil {
		log.Fatal(fcErr)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
