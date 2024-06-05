package main

import (
	"net/http"
)

func main() {
	Rotas()
	http.ListenAndServe(":8080", nil)
}
