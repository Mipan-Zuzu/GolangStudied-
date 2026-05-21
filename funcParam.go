package main
import "fmt"
func tambah (a int, b int) int {
	return a + b 
}

func param (name string, age int, loop []string) {
	fmt.Println("nama saya", name, "umur saya", age)

	for _, data := range loop {
		fmt.Println(data)
	}

}
