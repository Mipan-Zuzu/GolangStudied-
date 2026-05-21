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

func multiplevalue (name string, age int) (string, int) {
	return name, age
}

func namedRetunr () (fristname, secondname, thirdname string) {

	fristname = "mipan"
	secondname = "suzu"
	thirdname = "cuki"

	return fristname, secondname, thirdname
}

func lopfunc (data ...int) int {
	result := 0
	for _, a := range data {
		result += a
	}

	return result
}