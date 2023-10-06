package controllers

import (
	"html/template"
	"log"
	"main/models"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProtudos := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", todosOsProtudos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil{
			log.Panicln("Erro na conversao do preco:", err)
		}
		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil{
			log.Panicln("Erro na conversao da quantidade:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do ID para INT", err)
		}
		precoConvertidaParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do Preco para Float", err)
		}
		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade para int", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, quantidadeConvertidaParaInt, descricao, nome, precoConvertidaParaFloat, )

	}
	http.Redirect(w, r, "/", 301)
}
