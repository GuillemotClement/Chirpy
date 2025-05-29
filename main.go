package main

import (
	"log"
	"net/http"
)

func main() {
	// creation du constante pour le port utiliser
	const port = "8080"
	// definis une constante pour le dossier contenant les fichiers statique
	const filepathRoot = "./app"

	// creation du serveur
	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	// ajout d'une nouvelle route
	// on declare une fonction handler pour gerer la route
	mux.HandleFunc("/healthz", handleHealthz)

	// parametre du serveur avec la struct
	srv := http.Server{
		Handler: mux,        // definition du handler
		Addr:    ":" + port, // on concatene avec le port pour definir l'adresse
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe()) // pas besoin de passer le handler et addr qui sont definis dans la struct de parametre
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	// definir le header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	// definis le status code => ici 200
	w.WriteHeader(http.StatusOK)
	// ecriture du body
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
