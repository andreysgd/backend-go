package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		units := r.FormValue("units")

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Converting error:", err)
		}

		convertedUnits, err := strconv.Atoi(units)
		if err != nil {
			log.Println("Converting error:", err)
		}

		models.CreateNewProduct(name, description, convertedPrice, convertedUnits)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}
