package main

import "fmt"
import "log"

func withError() (string, error) {
	return "Привет", fmt.Errorf("%d ошибок", 10)
}

func main() {
	var name string
	var age int

	fmt.Print("Привет!\n")
	fmt.Print("Как тебя зовут?\n")
	fmt.Scanf("%s\n", &name)

	fmt.Print("Сколько тебе лет ?\n")
	fmt.Scanf("%d\n", &age)

	if age <= 0 {
		log.Fatal("Возраст должен быть > 0\n")
	}

	fmt.Printf("%s, тебе %d лет.", name, age)
	msg, err := withError()
	fmt.Println(msg)
	fmt.Println(err)
}
