package main

type Livro struct {
	Id      int
	Nome    string
	Autor   string
	Sinopse string
}

func BuscarProdutos() []Livro {

	db := ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Livro{}
	produtos := []Livro{}

	for selectDeTodosOsProdutos.Next() {
		var id int
		var nome_livro, autor_livro, sinopse_livro string

		err = selectDeTodosOsProdutos.Scan(&id, &nome_livro, &autor_livro, &sinopse_livro)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome_livro
		p.Autor = autor_livro
		p.Sinopse = sinopse_livro

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CreateProdict(nome, autor, sinopse string) {
	db := ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, autor, sinopse) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, autor, sinopse)
	defer db.Close()
}
func DeleteProduct(id string) {
	db := ConectaComBancoDeDados()

	delete, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Livro {
	db := ConectaComBancoDeDados()

	productDB, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productUpdate := Livro{}

	for productDB.Next() {
		var id int
		var nome, autor, sinopse string

		err = productDB.Scan(&id, &nome, &autor, &sinopse)

		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Nome = nome
		productUpdate.Autor = autor
		productUpdate.Sinopse = sinopse
	}

	defer db.Close()

	return productUpdate
}

func UpdateProduct(id int, nome, autor, sinopse string) {
	db := ConectaComBancoDeDados()

	updateProduct, err := db.Prepare("update produtos set nome=$1, autor=$2, sinopse=$3 where id=$4")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(nome, autor, sinopse, id)

	defer db.Close()
}
