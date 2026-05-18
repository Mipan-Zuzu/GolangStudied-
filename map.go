package main
import "fmt"

func maps () {
	user := map[string]string{  //TODO: Array
		"name" : "Mipan",
		"role" : "Admin",
	}

	user["name"] = "Sukis"
	delete(user, "name")

	// fmt.Println(user)
	fmt.Println(user["name"])

	// //TODO: ambil hanya name
	// fmt.Println(user["name"]) 

	// //TODO: Tambah data
	// user["age"] = "20"
	// user["role"] = "Owner"
	// user["hobbi"] = "Bersepeda"

	// _, nameOk := user["name"]
	// _, ageOk := user["age"];

	// fmt.Println(nameOk)

	// // Todo Cek value age kalo ok lanjut
	// if ageOk {
	// 	// TODO: Delete
	// 	delete(user, "age")
	// 	fmt.Println("berhasil terhapus age nya")
	// }


	// // TODO: LOOPING ARRAY
	// for key, value := range user {
	// 	fmt.Println( key , value)
	// }

}
