package main

import (
	"net/http"
	"strings"

	"github.com/unnxt30/Blog-Aggregator/internal/database"
)

type authedHandler func( http.ResponseWriter, *http.Request, database.User);

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract and validate the token
        header := r.Header.Get("Authorization")
        if header == "" {
            respondWithError(w, 401, "Authorization header is missing")
            return
        }
        parts := strings.Split(header, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            respondWithError(w, 401, "Invalid Authorization header format")
            return
        }
        token := parts[1]
        
        // Fetch user based on token
        user, err := cfg.DB.GetUserByApiKey(r.Context(), token)
        if err != nil {
            respondWithError(w, 401, "Invalid token")
            return
        }
        // Call the next handler with the authenticated user
        handler(w, r, user)
    }
}

