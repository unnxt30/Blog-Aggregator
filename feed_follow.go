package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/unnxt30/Blog-Aggregator/internal/database"
)


func (cfg *apiConfig) CreateFeedfollow(w http.ResponseWriter, r *http.Request, user database.User) {
    type FeedFollowStruct struct {
        FeedId string `json:"feed_id"`
    }

    // Read the body
    body, err := io.ReadAll(r.Body)
    if err != nil {
        respondWithError(w, 400, "problem reading request body")
        return
    }

    // Print the raw request body
    fmt.Println("Request Body:", string(body))

    // Create a new buffer with the body so we can read it again
    r.Body = io.NopCloser(bytes.NewBuffer(body))

    // Decode the JSON from the body
    var feed_id FeedFollowStruct
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&feed_id)
    if err != nil {
        respondWithError(w, 400, "problem creating Feed Follow")
        return
    }

    fmt.Printf("%v\n", feed_id.FeedId)
    feedID, _ := uuid.Parse(feed_id.FeedId)
    fmt.Printf("%v %v", feedID, user.ID)

    // DB operations and response

    FeedFollow, err := cfg.DB.CreateFeedfollow(r.Context(), database.CreateFeedfollowParams{
        FeedFollowID: uuid.New(),
        CreatedAt:    time.Now().Local().UTC(),
        UpdatedAt:    time.Now().Local().UTC(),
        FeedID:       feedID,
        UserID:       user.ID,
    })

    if err != nil {
        respondWithError(w, 401, "could not follow feed")
        return
    }

    respondWithJSON(w, 200, FeedFollow)
}


func (cfg *apiConfig) RemoveFeedfollow(w http.ResponseWriter, r *http.Request){
	removeID := r.PathValue("unfollowID");
	followFeedID, _ := uuid.Parse(removeID);
	cfg.DB.DeleteFeedfollow(r.Context(), followFeedID);
}