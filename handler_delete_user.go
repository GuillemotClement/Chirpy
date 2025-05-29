package main

import (
	"net/http"
	"os"
)

func (cfg *apiConfig) HandlerResetUser(w http.ResponseWriter, r *http.Request) {
	platform := os.Getenv("PLATFORM")

	if platform != "dev" {
		respondEndpointForbidden(w, http.StatusForbidden, "Endpoint forbidden")
		return
	}

	err := cfg.db.DeleteAllUser(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed delete user", err)
		return
	}
	res := map[string]string{
		"status": "Delete user is done",
	}
	respondWithJSON(w, http.StatusOK, res)
}
