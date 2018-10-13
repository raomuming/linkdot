package main

import (
	"github.com/raomuming/linkdot/route"
	"log"
	"net/http"
)

func main() {
	r := route.NewRouter()

	log.Println("start listening port 8080")
	http.ListenAndServe(":8080", r)
}
