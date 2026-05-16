package main

import (
	"fmt"
)

func switchs() {
	var grade = "senin"

	switch grade {
	case "senin":
		fmt.Println("Gawe")
	case "minggu":
		fmt.Println("Libur gwe")
	default:
		fmt.Println("anjay")
	}
}

func loopings () {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("p")
		if (i == 5) {
			fmt.Println("cukuppp")
			break
		}
	}
}
