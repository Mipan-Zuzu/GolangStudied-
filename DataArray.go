package main
import "fmt"

func Arry () {
	var names [3] string
	names[0] = "mipan"
	names[1] = "Cuki"
	names[2] = "Liars"

	for _, name := range names {
		fmt.Println(name)
	}
}