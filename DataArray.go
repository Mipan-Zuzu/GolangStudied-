package main
import "fmt"

var values = [3] int {
	30, 40, 50,
}

func Arry () {
	var names [3] string
	names[0] = "mipan"
	names[1] = "Cuki"
	names[2] = "Liars"

	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println("ini values", values)
	Slices()
}

func Slices () {
	namess := [...]string{"mipan", "zuzu", "cuki", "liar", "rrampe", "Bawdewak"}
	slice := namess[4:] //slice

	names2 := append(slice, "cukiliar")
	names3 := len(names2)

	names4 := make([]string, 10)
	names5 := append(names2, "Woilah")
	names6 := append(names4, "Ikan goreng")
	names7 := copy(names4, names2)

	fmt.Println(slice[0])
	fmt.Println(names2)
	fmt.Println(names3)
	fmt.Println("ini adalah hasil dari make" , names4)
	fmt.Println(names5)
	fmt.Println(names6, len(names4))
	fmt.Println(names7)
}