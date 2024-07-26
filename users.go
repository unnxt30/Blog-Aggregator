package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/unnxt30/Blog-Aggregator/internal/database"
)

type UsrStruct struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name string `json:"name"`

}

// func dbUsertoUsrStruct(dbUser database.CreateUserParams) UsrStruct{

// 	return UsrStruct{
// 		ID: dbUser.ID,
// 		CreatedAt: dbUser.CreatedAt,
// 		UpdatedAt: dbUser.UpdatedAt,
// 		Name: dbUser.Name.String,
// 	}

// }

func (cfg *apiConfig)HandleUserFunc(w http.ResponseWriter, r *http.Request){
	type UserRequest struct{
		Name string 
	}
	decoder := json.NewDecoder(r.Body);
	userName := UserRequest{};
	
	err := decoder.Decode(&userName);
	if err != nil{
		respondWithError(w, 500, "error creating user");
	}
	
	createdUser, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name: sql.NullString{String: userName.Name, Valid: true},
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
		ID: uuid.New(),
	}) 
	

	if err != nil {
		respondWithError(w, 501, "error creating a user");
	}

	respondWithJSON(w, 200, createdUser);
	};

func (cfg *apiConfig)GetUser(w http.ResponseWriter, r *http.Request){
	
	header := r.Header.Get("Authorization: ");
	headerArgs := strings.TrimPrefix(header, "ApiKey");
	if len(headerArgs) < 2{
		respondWithError(w, 401, "Invalid Header");
	}

	returnUser, err := cfg.DB.GetUserByApiKey(r.Context(), headerArgs);
	
	if err != nil{
		respondWithError(w, 400, "Bad Request")
	}

	respondWithJSON(w, 200, returnUser);

}
	
