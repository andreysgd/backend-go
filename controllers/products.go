package controllers

import (
	"net/http"
	"text/template"

	"github.com/web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
