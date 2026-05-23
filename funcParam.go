package main
import "fmt"
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

func goodbay (name string)string {
	return "goodbay goofy ah" + name
}


func checking (data int,tampilkan func(string) string)string {
	if data %2 == 0 {
	 return tampilkan("genap")
	}

	return tampilkan("ganjil")
}

func tampilkan (data string) string {
	return data
}

type GenreType func(string) bool

func chekingGenre (genre string, genreType GenreType) {
	if genreType(genre) {
		fmt.Println("dilarang")
		return
	}

	fmt.Println("boleh coy")
}

func factorialRecrusive (value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}
