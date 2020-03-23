package main

import (
	"net/http"

	"github.com/himynameisoleg/simple-webservice-go/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
