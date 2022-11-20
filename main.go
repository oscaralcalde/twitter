package main

import (
	"log"

	"github.com/oscaralcalde/twitter/bd"
	"github.com/oscaralcalde/twitter/handlers"
)

func main() {
	if bd.ConnectionCheck() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Handlers()
}
