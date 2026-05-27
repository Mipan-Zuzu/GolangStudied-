package main

import (
	"fmt"
	"net/http"
)

func handlers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`{"data" : "pong", "status" : 200 }`))
	fmt.Println("succsesfuly")
}

