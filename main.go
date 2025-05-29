package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"Github.com/GuillemotClement/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
}

func main() {
	// chargement du fichier .env
	godotenv.Load()
	// recuperation de la db url
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	// connection db
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	const port = "8080"
	const filepathRoot = "./app"

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries, // on rajoute la connection dans la struct
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))) // middleware permet de compter le nombre de request recus
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	// methode de la struct. Il faut faire reference a cette struct pour acceder a la methode
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)

	srv := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
