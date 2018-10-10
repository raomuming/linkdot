package main

import (
	"github.com/raomuming/linkdot/route"
	"net/http"
)

func main() {
	r := route.NewRouter()

	http.ListenAndServe(":8080", r)
}
