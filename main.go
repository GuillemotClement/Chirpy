package main

import (
	"log"
	"net/http"
)

func main() {
	// définition du port dans une variable
	const port = "8080"

	// création du nouveau routeur
	mux := http.NewServeMux()

	// configuration du serveur
	srv := &http.Server{
		// addresse et port écouter par le serveur
		Addr: ":" + port,
		// définis le routeur qui reçoit les requêtes
		Handler: mux,
	}
	log.Printf("Serveur is listen on port : %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
