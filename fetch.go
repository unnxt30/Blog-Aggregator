package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/unnxt30/Blog-Aggregator/internal/database"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Atom    string   `xml:"atom,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Guid        string `xml:"guid"`
			Description string `xml:"description"`
		} `xml:"item"`
	} `xml:"channel"`
} 

func (cfg *apiConfig) feedWorker(ctx context.Context){
	feedArr, err := cfg.DB.GetNextFeedsToFetch(ctx, 10)
	if err != nil {
		// Handle error fetching feeds
		fmt.Println("Error fetching feeds:", err)
		return
	}

	var wg sync.WaitGroup

	for _, feed := range feedArr {
		// Mark feed as fetched before processing
		err = cfg.DB.MarkFeedFetched(ctx, feed.FeedID)
		if err != nil {
			// Handle error marking feed as fetched
			fmt.Println("Error marking feed as fetched:", err)
			continue // Skip this feed and proceed
		}
	}

	for _, feed := range feedArr {
		wg.Add(1)

		// Capture feed for each goroutine
		go func(feed database.Feed) {
			defer wg.Done()

			feedData := FetchFeedData(string(feed.Url.String))
			
			for _, item := range feedData.Channel.Item {
				// Additional logging
				fmt.Println("Fetched title:", item.Title)
			}
		}(feed) // Passing feed as an argument to capture its value
	}

	wg.Wait()
}


func FetchFeedData(url string) Rss {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(error.Error(err))
		return Rss{}
	}
	defer resp.Body.Close() // Ensures body closes even if there's a read/unmarshal error

	xmlData := Rss{}
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println(err.Error())
		return Rss{}
	}

	err = xml.Unmarshal(body, &xmlData) // Unmarshalling directly from body
	if err != nil {
		fmt.Println(err.Error())
	}

	return xmlData
}

func (cfg *apiConfig) FetchFeeds(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure context cancellation when main exits

	// Create a ticker that ticks every 60 seconds
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	// Run the worker initially before starting the ticker
	go cfg.feedWorker(ctx)

	// Start a loop that will call feedWorker every 60 seconds
	for {
		select {
		case <-ticker.C:
			go cfg.feedWorker(ctx)
		// Handle context cancellation
		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping the worker")
			return
		}
}

}