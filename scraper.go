package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/ShubhamTiwari55/helloGo/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
){
	log.Printf("Scarping started with %v goroutines every %v duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextfeedToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {	
			log.Printf("Error getting feeds to fetch: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scraperFeed(wg, feed, db)
		}
		wg.Wait()
	}
}

func scraperFeed(wg *sync.WaitGroup, feed database.Feed, db *database.Queries){
	defer wg.Done()
	//scraping the feed

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error getting feed: %v", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing time: %v", err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title: item.Title,
			Description: description,
			PublishedAt: t,
			Url: item.Link,
			FeedID: feed.ID,
	})
	if err != nil {
		log.Printf("Error creating post: %v", err)
		continue
	}

	log.Printf("Feed %s collected, %d posts", feed.Name, len(rssFeed.Channel.Item))
}}