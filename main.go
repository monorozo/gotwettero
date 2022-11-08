package main

import (
	"log"

	"github.com/monorozo/gotwettero/bd"
	"github.com/monorozo/gotwettero/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadore()
}
