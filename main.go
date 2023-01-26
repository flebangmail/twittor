package main

import (
	"log"
	"github.com/flebangmail/CursoGo/bd"
	"github.com/flebangmail/CursoGo/handlers"
)

func main() {
	if !bd.ChequeoConexion() {
		log.Fatal("Sin Conexion a la BD") 
	}
	handlers.Manejadores()
}