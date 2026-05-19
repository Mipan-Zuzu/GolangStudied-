package main
import "fmt"

func param (name string, age int, loop []string) {
	fmt.Println("nama saya", name, "umur saya", age)

	for _, data := range loop {
		fmt.Println(data)
	}
}