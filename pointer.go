package main
import "fmt"

type Adress struct {
	City, Province, Country string
}

func pointers () {
	addres1 := Adress{"Denpasar", "Bali", "indoneisa"}
	addres2 := &addres1

	addres2.City = "SukaWati"

	fmt.Println(addres1)
	fmt.Println(addres2)
}