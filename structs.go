package main

import "fmt"


// TODO: mirip di typecript interface
type User struct {
	Name string
	Age  int
	adress string
}

func structs() {
	user := User{
		Name: "mipan",
		Age:  30,
		adress : "jalan mawar",
	}

	fmt.Println(user.Name)
	user.Age = 21
	fmt.Println(user)
}
