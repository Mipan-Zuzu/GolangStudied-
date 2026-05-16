package main
import "fmt"

type nums int

func BindingOperation () {
	const a nums = 10
	const b nums = 10

	const name string = "ilham"
	const gradeIPA = 80
	const gradeMTK = 50

	const resultIPA = gradeIPA >= 75 
	const resultMTK = gradeMTK >= 80
	
	var checkGrade =  resultIPA && resultMTK
	fmt.Println(checkGrade)
}