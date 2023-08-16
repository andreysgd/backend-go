package main

import (
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Name        string
	Description string
	Price       float64
	Units       int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Name: "Dell G15", Description: "Nice hardware", Price: 4499, Units: 123},
		{"Acer Nitro 5", "Perfect for Gamers", 5699, 57},
		{"Alienware V13", "Behold inovation", 12899, 23},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
