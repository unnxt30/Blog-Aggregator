package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/unnxt30/Blog-Aggregator/internal/database"
)

type FeedParam struct{
	FeedID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name string `json:"name"`
	Url string `json:"url"`
	UserID uuid.UUID `json:"user_id"`

}


func simplifyFeedParam(params database.Feed) FeedParam{
	return FeedParam{
		FeedID: params.FeedID,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
		Name: params.Name.String,
		Url: params.Url.String,
		UserID: params.UserID,
	}
}

func (cfg *apiConfig) PostFeed(w http.ResponseWriter, r *http.Request, usr database.User){
	type feedParam struct{
		Name string
		Url string
	};
	
	var recievedData feedParam;
	decoder := json.NewDecoder(r.Body);
	err := decoder.Decode(&recievedData);
	if err != nil{
		respondWithError(w, 400, "Invalid Body")
	}

	returnFeed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		FeedID: uuid.New(),	
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
		Name: sql.NullString{String: recievedData.Name, Valid: true},
		UserID: usr.ID,
		Url: sql.NullString{String: recievedData.Url, Valid: true},
	})

	respondWithJSON(w, 200, simplifyFeedParam(returnFeed));

		
}


