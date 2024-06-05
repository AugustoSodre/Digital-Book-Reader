package main

import (
	"net/http"
)

func Rotas() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/criar-produto", CreateProduct)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/update", Update)

}
