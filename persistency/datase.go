package persistency

import (
	"Api-Aula1/config"
	"database/sql"
	"log"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.Cfg.FormatDSN())
	if err != nil {
		log.Println("Erro ao abrir conexão com o DB")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Println("Erro ao pingar no DB")
		return nil, err
	}

	return db, nil
}
