package main

import (
	"log"
	"net/http"
)

func main() {
	// creation du constante pour le port utiliser
	const port = "8080"
	// creation du serveur
	mux := http.NewServeMux()
	// parametre du serveur avec la struct
	srv := http.Server{
		Handler: mux,        // definition du handler
		Addr:    ":" + port, // on concatene avec le port pour definir l'adresse
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(srv.ListenAndServe()) // pas besoin de passer le handler et addr qui sont definis dans la struct de parametre
}
