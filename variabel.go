package main

import (
	"fmt"
)

var name = "mipan"
var ages = 40

var price float64 
var discount float64
var grade byte = 'A'

var endcode byte = 'M'

func sayName (food string, hobbies string, married bool) {
	fmt.Println(name, ages, food, hobbies, married)
}

func CountNum () {
	price = 44.2
	discount = 8.4
	fmt.Println(price - discount)
}

func Grades () {
	println(grade)
	println(string(grade))
}

func Emoji () {
	var latter rune
	latter = '😅'
	println(latter)
	println(string(latter))
}