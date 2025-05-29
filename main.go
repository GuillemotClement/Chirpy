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
	platform       string
}

func main() {
	const port = "8080"
	const filepathRoot = "./app"

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatalf("PLATFORM must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries, // on rajoute la connection dans la struct
		platform:       platform,
	}

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))) // middleware permet de compter le nombre de request recus
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	// methode de la struct. Il faut faire reference a cette struct pour acceder a la methode
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerResetUser)
	mux.HandleFunc("POST /api/validate_chirp", handlerChirpsValidate)
	// creation user
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)

	srv := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
