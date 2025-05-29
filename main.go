package main

import (
	"log"
	"net/http"
)

func main() {
	// creation du constante pour le port utiliser
	const port = "8080"
	// definis une constante pour le dossier contenant les fichiers statique
	const filepathRoot = "."

	// creation du serveur
	mux := http.NewServeMux()

	// ajout d'une nouvelle route
	// on indique le folder root comme contenant le fichier html
	// retourne automatiquement le fichier index.html
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	// parametre du serveur avec la struct
	srv := http.Server{
		Handler: mux,        // definition du handler
		Addr:    ":" + port, // on concatene avec le port pour definir l'adresse
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe()) // pas besoin de passer le handler et addr qui sont definis dans la struct de parametre
}
