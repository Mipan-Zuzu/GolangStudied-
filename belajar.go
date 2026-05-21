package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

var food string
var hobbies = "gowes"

func main() {
	// married := false
	// food = "burger"

	// if (food == "burger") {
	// 	fmt.Println("iya saya suka berger")
	// }else {
	// 	fmt.Println("saya gak suka berger")
	// }

	// if(married != true) {
	// 	fmt.Println("aku belum menikah")
	// }

	fmt.Println(quote.Go())
	// sayName(food, hobbies, married)

	// maps()
	// structs()
	// pointers()

	var loop = []string{"finding", "found", "love", "gone", "finding Again"}

	param("Mipan", 17, loop)
	fmt.Println(tambah(3, 5))
	names, _ := multiplevalue("mipan", 20)

	fmt.Println(names)

	fristname, secondname,_ := namedRetunr()

	fmt.Println(fristname, secondname)
	result := lopfunc(1,4,5,7,98)
	fmt.Println(result)
}
