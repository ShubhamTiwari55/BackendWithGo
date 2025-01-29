package main

import (
	"fmt"
	"net/http"

	"github.com/ShubhamTiwari55/helloGo/internal/auth"
	"github.com/ShubhamTiwari55/helloGo/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	handler(w, r, user)
	}
}