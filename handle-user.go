package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ShubhamTiwari55/helloGo/internal/auth"
	"github.com/ShubhamTiwari55/helloGo/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:name`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing json:", err))
		return
	}
	
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
})

if err!= nil {
	respondWithError(w, 400, fmt.Sprint("Error creating user: ", err))
	return
}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error getting api key: ", err))
		return
	}
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error getting user: ", err))
		return
	}
	respondWithJSON(w, 200, databaseUserToUser(user))
}