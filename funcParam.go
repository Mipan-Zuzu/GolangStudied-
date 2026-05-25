package main

import (
	"errors"
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



type Animal interface {
	Sound()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) Sound() {
	fmt.Println(c.Name, "meong")
}

func (d Dog) Sound() {
	fmt.Println(d.Name, "Woff")
}

func Test (a Animal) {
	a.Sound()
}



func TambahCuayo (data interface{} )error {
	value, typeVal := data.(int)
	if !typeVal {
		fmt.Println("data bukan integer")
		return errors.New("type must int")
	}
	
	fmt.Println(value + 5)
	return nil
}

func newMap (data string) map[string] string {
	if data == "" {
		return nil
	}

	return map[string]string {
		"name" : data,
	}
}

func TypeAssort () interface{} {
	return 77
}

type Alamat struct {
	name , lokasi string 
	age int
}

func Pointah () {
	data := Alamat{"Mipan", "Bali", 17}
	data2 := &data

	data2.name = "Cuki"
	fmt.Println(data)
	fmt.Println(data2)
}

func asteriskOp (data int) {
	result := data + 4
	count := &result

	fmt.Println(*count)
}