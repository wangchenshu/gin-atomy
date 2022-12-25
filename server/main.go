package main

import (
	"log"
	"net/http"

	"gin-atomy/server/routes"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", routes.Engine()))
}
