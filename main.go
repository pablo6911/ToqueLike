package main

import (
	"log"

	"github.com/pablo6911/ToqueLike/bd"
	"github.com/pablo6911/ToqueLike/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
