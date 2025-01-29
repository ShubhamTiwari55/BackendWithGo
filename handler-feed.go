package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/ShubhamTiwari55/helloGo/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:name`
		URL string `json:url`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing json:", err))
		return
	}
	
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		UserID: user.ID,
		Url: params.URL,
})

if err!= nil {
	respondWithError(w, 400, fmt.Sprint("Error creating feed: ", err))
	return
}

	respondWithJSON(w, 200, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

if err!= nil {
	respondWithError(w, 400, fmt.Sprint("Couldn't get feeds: ", err))
	return
}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}