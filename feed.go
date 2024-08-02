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
		return
	}

	returnFeed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		FeedID: uuid.New(),	
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
		Name: sql.NullString{String: recievedData.Name, Valid: true},
		UserID: usr.ID,
		Url: sql.NullString{String: recievedData.Url, Valid: true},
	})
	
	if err != nil{
		respondWithError(w, 400, "could not create feed");
		return;
	}

	FeedFollow, err := cfg.DB.CreateFeedfollow(r.Context(), database.CreateFeedfollowParams{
        FeedFollowID: uuid.New(),
        CreatedAt:    time.Now().Local().UTC(),
        UpdatedAt:    time.Now().Local().UTC(),
        FeedID:       returnFeed.FeedID,
        UserID:       usr.ID,
    })

	if err != nil {
		respondWithError(w, 400, "could not follow feed")
		return;
	}
	respondWithJSON(w, 200, map[string]interface{}{"feed": simplifyFeedParam(returnFeed), "feed_follow": FeedFollow});


		
}

func (cfg *apiConfig) GetAllFeed(w http.ResponseWriter, r *http.Request){
	feeds, err := cfg.DB.GetFeed(r.Context());

	if err != nil{
		respondWithError(w, 400, "Could not get feed");
	}

	respondWithJSON(w, 200, feeds);

}


