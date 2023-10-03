package models

import (
	"main/db"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
	Id              int
}

func SearchProducts() []Produto {
	db := db.ConnectionDB()

	selectDeTodosOsProdutos, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectionDB()

	insertDados, err := db.Prepare("insert into products(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
