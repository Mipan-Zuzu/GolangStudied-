package main
import "fmt"

func perulangan () {
	counter := "masuk"

	name := [...]string{"rice", "food", "home"}

	for i , counters := range counter {
		fmt.Println(i, counters)
	}

	for i, data := range name {
		fmt.Println("data num", i, data)
	}
}