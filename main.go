package main

import (
	"Api-Aula1/config"
	"Api-Aula1/router"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	r := router.New()
	log.Printf("Servidor ouvindo em %s ...", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r))
}
