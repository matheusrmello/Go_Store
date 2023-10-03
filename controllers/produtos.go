package controllers

import (
	"html/template"
	"main/models"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProtudos := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", todosOsProtudos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
