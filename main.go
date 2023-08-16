package main

import (
	"net/http"

	"github.com/web-application/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
