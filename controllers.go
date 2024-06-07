package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := BuscarProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "criar-produto", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		autor := r.FormValue("autor")
		sinopse := r.FormValue("sinopse")

		CreateProdict(nome, autor, sinopse)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	DeleteProduct(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := EditProduct(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		autor := r.FormValue("autor")
		sinopse := r.FormValue("sinopse")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter o Id para int: ", err)
		}

		UpdateProduct(idConv, nome, autor, sinopse)

	}

	http.Redirect(w, r, "/", 301)
}
