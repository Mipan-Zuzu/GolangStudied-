package main
import "fmt"

func Switchs () {
	// day := "sunday"
	isDay := "ini"
	var leng int = len(isDay)

	switch  {
	case leng > 5 :
		fmt.Println("ini minggu")
	case leng < 5:
		fmt.Println("ini hari senin")
	}
}