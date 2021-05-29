package main

import (
	"github.com/asaldana-uoc/tfg-uoc-app/pkg/handlers"
	"log"
	"net/http"
)

const port = ":8888"

// Main principal d'entrada de l'aplicació on es defineix el port on s'aixecarà el servidor web,
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/estat", handlers.Status)
	log.Printf("Aplicació Web de test del TFG iniciada i escoltant al port %s", port)
	_ = http.ListenAndServe(port, nil)
}



