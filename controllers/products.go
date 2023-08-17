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
	httpStatusCode := 301
	http.Redirect(w, r, "/", httpStatusCode)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	httpStatusCode := 301
	http.Redirect(w, r, "/", httpStatusCode)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		units := r.FormValue("units")

		convertedId, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Converting error:", err)
		}

		convertedPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Converting error:", err)
		}

		convertedUnits, err := strconv.Atoi(units)
		if err != nil {
			log.Println("Converting error:", err)
		}

		models.UpdateProduct(convertedId, name, description, convertedPrice, convertedUnits)
	}
	httpStatusCode := 301
	http.Redirect(w, r, "/", httpStatusCode)
}
