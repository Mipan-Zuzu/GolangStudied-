package main

import (
	"fmt"

)
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

func loggin () {
	fmt.Println(" 1. User melakukan login")
}

func authorize () {
	defer loggin()
	fmt.Println(" 2. authorization user ")
}

func succses () {
	authorize()
	fmt.Println("3. succses login")
}

func endApp () {
	fmt.Println("aplikasi di hentikan")
	message := recover()
	fmt.Println("terjadi error di ", message)
}

func runApp (err bool) {
	defer endApp()
	if err {
		panic("ERROR")
	}
}
func sistem () {
	fmt.Println("sistem jalan")
}

type Users struct {
	name, adres string
	age int
	maried bool
}

func callUser () {
	var Mipan Users
	Mipan.name = "Mipan"
	Mipan.adres = "bali"
	Mipan.age = 17
	Mipan.maried = true

	fmt.Println(Mipan)
}

func Buyer () {
	Mipan := Users {
		name: "CukiMipan",
		adres: "bern",
		age: 17,
		maried: true,
	}
	fmt.Println(Mipan)

	Cuki := Users{"Cuki", "syedney", 17, true}
	fmt.Println(Cuki)
}

type Customer struct {
	name string
}

func (customer Customer) Greating() {
	fmt.Println("welcome customer", customer.name)
}

type Change struct {
	name string
	money int
}

func (change Change) sayChange() {
	barangPrice := 10
	if change.money <= barangPrice {
		fmt.Println("maaf uang tidak cukup ", change.name)
		return
	}
		fmt.Println("terimakasi kembali ", change.name)
}