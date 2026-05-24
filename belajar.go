package main

import (
	"fmt"

	"rsc.io/quote/v4"
)

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

	// var loop = []string{"finding", "found", "love", "gone", "finding Again"}

	// param("Mipan", 17, loop)
	// fmt.Println(tambah(3, 5))
	// names, _ := multiplevalue("mipan", 20)

	// fmt.Println(names)

	// fristname, secondname,_ := namedRetunr()

	// fmt.Println(fristname, secondname)
	// result := lopfunc(1,4,5,7,98)
	// fmt.Println(result)

	// count := []int{22,44,33}
	// totals := lopfunc(count...)
	// fmt.Println(totals)
	// bay := goodbay
	// fmt.Println(bay("mipan"))
	

	// genre := func (name string) bool {
	// 	return name == "romance"
	// }

	// chekingGenre("sliceoflife", genre)

	// fmt.Println(factorialRecrusive(5))

	// count := 0

	// startCount := func () {
	// 	for i := 0 ; i < 10; i ++ {
	// 		count += i
	// 	}
	// }

	// startCount()
	// fmt.Println(count)


	// succses()
	// runApp(true)
	// callUser()
	// Buyer()

	// Mipan := Customer{name: "Miapn"}
	// Cuki := Customer{name: "Cuki"}
	// Mipan.Greating()
	// Cuki.Greating()
	
	// MipanPok := Change{"MipanPok3", 9}
	// MipanPok.sayChange()
	
	// kucing := Cat{Name:  "Cuki"}
	// anjing := Dog{Name:  "AmbaCenat"}

	// Test(kucing)
	// Test(anjing)

	// cuayo := TambahCuayo("dasda")
	// fmt.Println(cuayo)

	// mapResult := newMap("")
	// fmt.Println(mapResult)

	// data := newMap("")
	// if data == nil {
	// 	fmt.Println("data map kosong")
	// }
	// fmt.Println(data)

	ssort := TypeAssort()
	ssortString := ssort.(string)
	fmt.Println(ssortString)

	ssortInt := ssort.(int)
	fmt.Println(ssortInt)
	
}
