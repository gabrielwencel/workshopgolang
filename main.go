package main

import (
	"Api-Aula1/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()

	const addr = ":8080"
	log.Printf("Servidor ouvindo em %s ...", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
