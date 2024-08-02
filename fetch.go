package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/unnxt30/Blog-Aggregator/internal/database"
)

const layout = "Mon, 02 Jan 2006 15:04:06 -0700";

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
	
	feedArr, err := cfg.DB.GetNextFeedsToFetch(ctx, 1)
	if err != nil {
		// Handle error fetching feeds
		fmt.Println("Error fetching feeds:", err)
		return
	}

	var wg sync.WaitGroup

	for _, feed := range feedArr {
		err = cfg.DB.MarkFeedFetched(ctx, feed.FeedID)
		if err != nil {
			fmt.Println("Error marking feed as fetched:", err)
			continue
		}
	}

	for _, feed := range feedArr {
		wg.Add(1)

		go func(feed database.Feed) {
			defer wg.Done()
			
			feedData := FetchFeedData(string(feed.Url.String))
			
			for _, item := range feedData.Channel.Item {
				parsedTime, err := time.Parse(layout, item.PubDate)
				cfg.DB.CreatePost(ctx, database.CreatePostParams{
					ID: uuid.New(),
					CreatedAt: time.Now().Local().UTC(),
					UpdatedAt: time.Now().Local().UTC(),
					Title: item.Title,
					Url: item.Link,
					PublishedAt: sql.NullTime{Time: parsedTime, Valid: true},
					FeedID: feed.FeedID,
				});
				


				if err != nil {
					fmt.Println(err)
					continue;
				}

				fmt.Println("Fetched title date:", parsedTime);
			}
		}(feed) 
	}
	wg.Wait()
}


func FetchFeedData(url string) Rss {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(error.Error(err))
		return Rss{}
	}
	defer resp.Body.Close() 

	xmlData := Rss{}
	body, err := io.ReadAll(resp.Body)
	
	if err != nil {
		fmt.Println(err.Error())
		return Rss{}
	}

	err = xml.Unmarshal(body, &xmlData) 
	if err != nil {
		fmt.Println(err.Error())
	}

	return xmlData
}

func (cfg *apiConfig) FetchFeeds(){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() 

	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	go cfg.feedWorker(ctx)
	
	for {
		select {
		case <-ticker.C:
			go cfg.feedWorker(ctx)
		case <-ctx.Done():
			fmt.Println("Context cancelled, stopping the worker")
			return
		}
	}
}