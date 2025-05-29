package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	// struct pour le json recus
	type parameters struct {
		Body string `json:"body"`
	}
	// struct pour le json renvoyer
	type returnVals struct {
		Cleaned_body string `json:"cleaned_body"`
	}

	// on recupere le body de la request
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// verification de la longueur du chirps
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "chirp is too long", nil)
		return
	}

	// creation du map des mauvais mot
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	// appelle de la fonction pour nettoyer le chirps
	cleaned := getCleanedBody(params.Body, badWords)

	// on retourne le chirps nettoyer
	respondWithJSON(w, http.StatusOK, returnVals{
		Cleaned_body: cleaned,
	})
}

// fonction qui nettoie le chirps
// porends une string
// prend un map des bad words
func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ") // on tranforme en liste
	for i, word := range words {
		// on parcour les mots du chirps
		// on les passe en lower
		loweredWord := strings.ToLower(word)
		// si un mot match avec un bad word
		if _, ok := badWords[loweredWord]; ok {
			// on remplace le mot
			words[i] = "****"
		}
	}
	// on rassemble les elemnts de la liste dans une string nettoyer
	cleaned := strings.Join(words, " ")
	return cleaned
}
