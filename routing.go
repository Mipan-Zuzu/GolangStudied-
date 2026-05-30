package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handlers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(`{"data" : "pong", "status" : 200 }`))
	fmt.Println("succsesfuly")
}

type Siswa struct {
	Name, School string
	Age, Grade   int
}

func dataSiswa(res http.ResponseWriter, req *http.Request) {
	result := Siswa{
		Name:   "mipan",
		School: "smk 1 sukawati",
		Age:    17,
		Grade:  3,
	}

	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(result)
}

var siswslist []Siswa

type data struct {
	Id   int
	Data []Siswa
}

func idSiswa(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/api/user/")
	result, _ := strconv.Atoi(id)
	res.Header().Set("Content-Type", "application/json")
	data := data{
		Id:   result,
		Data: siswslist,
	}
	json.NewEncoder(res).Encode(data)
}

type resultSave struct {
	Status int
	Data   Siswa
}

func SaveSiswa(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

	res.WriteHeader(200)
	var siswa Siswa
	res.Header().Set("Content-Type", "application/json")
	json.NewDecoder(req.Body).Decode(&siswa)
	siswslist = append(siswslist, siswa)
	var data = resultSave{
		Status: 200,
		Data:   siswa,
	}
	json.NewEncoder(res).Encode(data)
}

type errorType struct {
	Message string
	Status  int
}

type EditResult struct {
	Status int
	data   Siswa
}

func updateSiswa(res http.ResponseWriter, req *http.Request) {
	errorMessage := errorType{
		Message: "Wrong http req Method",
		Status:  405,
	}
	if req.Method != "PATCH" {
		res.WriteHeader(405)
		json.NewEncoder(res).Encode(errorMessage)
		return
	}
	id := strings.TrimPrefix(req.URL.Path, "/api/user/")
	if id == "" {
		res.WriteHeader(404)
		return
	}

	var siswa Siswa
	json.NewDecoder(req.Body).Decode(&siswa)
	result, _ := strconv.Atoi(id)
	siswslist[result] = siswa
	data := resultSave{
		Status: 200,
		Data:   siswa,
	}
	json.NewEncoder(res).Encode(data)
}
	