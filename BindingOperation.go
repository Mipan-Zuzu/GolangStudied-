package main
import "fmt"

type nums int

func BindingOperation () {
	const a nums = 10
	const b nums = 10

	var result bool = a < b
	fmt.Println(result)
}