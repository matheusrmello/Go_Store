package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectionDB() *sql.DB {
	conexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"0.0.0.0", 5051, "postgres", "qwerty", "alura_db")
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}
	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade      int
	Id              int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    db := connectionDB()

    selectDeTodosOsProdutos, err := db.Query("select * from produtos")
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

    temp.ExecuteTemplate(w, "Index", produtos)
    defer db.Close()
}
