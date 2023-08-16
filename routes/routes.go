package routes

import (
	"net/http"

	"github.com/web-application/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
}
