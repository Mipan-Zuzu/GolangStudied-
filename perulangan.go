package main
import "fmt"

func perulangan () {
	counter := "masuk"
	var numbers = []int{10, 44, 50}

	num := 11

	name := [...]string{"rice", "food", "home"}

	for i , counters := range counter {
		fmt.Println(i, counters)
	}

	for i, data := range name {
		fmt.Println("data num", i, data)
	}

	for _, data := range numbers {
		if(data < 9) {
			break
		}

		fmt.Println("now the data was", data)
	}


	for i := 0; i < num; i++ {
		if(i > 9) {
			break
		}
		fmt.Println("now", i)
	}

	for i := 0; i < num; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("sekarang data ke", i)
	}
}