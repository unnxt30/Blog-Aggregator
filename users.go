package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
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

func dbUsertoUsrStruct(dbUser database.CreateUserParams) UsrStruct{

	return UsrStruct{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name.String,
	}

}

func (cfg *apiConfig)HandleUserFunc(w http.ResponseWriter, r *http.Request){
	type UserRequest struct{
		Name string 
	}
	decoder := json.NewDecoder(r.Body);
	userName := UserRequest{};
	
	// fmt.Println(decoder);
	// fmt.Println("Hello")

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

	respondWithJSON(w, 200, dbUsertoUsrStruct(database.CreateUserParams(createdUser)));
	};


	
