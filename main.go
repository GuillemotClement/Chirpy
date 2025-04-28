package main

import (
	"log"
	"net/http"
)

func main() {
	// définition du port dans une variable
	const port = "8080"
	// définition du folder contenant les fichiers statiques
	const filePathRoot = "."

	// création du nouveau routeur
	mux := http.NewServeMux()

	// configuration du serveur
	srv := &http.Server{
		// addresse et port écouter par le serveur
		Addr: ":" + port,
		// définis le routeur qui reçoit les requêtes
		Handler: mux,
	}

	// pour la route "/", on viens retourner le fichier index.html
	// placer dans le root du projet
	// on pointe sur le serveurMux pour ajouter un handler qui gère le /
	mux.Handle("/", http.FileServer(http.Dir(filePathRoot)))

	// affichage dans le terminal
	log.Printf("Serveur is listen on port : %s\n", port)
	// affiche et stop si une erreur est lancer
	log.Fatal(srv.ListenAndServe())
}
