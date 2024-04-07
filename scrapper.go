package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/esiebomaj/rssarg/internal/database"
	"github.com/google/uuid"
)

func ScrapeFeeds(db *database.Queries, DurationInMins int16, concurency int16) {
	duration := time.Duration(DurationInMins) * time.Second
	ticker := time.NewTicker(duration)

	for range ticker.C {
		FetchFeeds(db, concurency)
	}

}

func FetchFeeds(db *database.Queries, limit int16) {
	// fetch limit feeds for db
	fmt.Println("Fetching Feeds from db")
	feeds, err := db.GetEarliestFetchedFeeds(context.Background(), int32(limit))
	if err != nil {
		fmt.Println("err fetching feeds from db", err)
		return
	}
	fmt.Println(feeds)

	wg := sync.WaitGroup{}
	for _, feed := range feeds {
		wg.Add(1)
		go FetchFeed(db, feed, &wg)
	}

	wg.Wait()
}

func FetchFeed(db *database.Queries, DBFeed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	// fetch posts for this feed
	rssFeed, err := URLToRSS(DBFeed.Url)
	if err != nil {
		fmt.Printf("Error fetching rss for %v : %v ", DBFeed.Url, err)
	}

	// save posts in db
	fmt.Println(rssFeed)
	for _, post := range rssFeed.Channel.Items {
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		PubDate, err := time.Parse(layout, post.PubDate)
		if err != nil {
			fmt.Printf("Cannot parse time %v for post %v", post.PubDate, post.Title)
			continue
		}
		db.CreatePost(context.Background(), database.CreatePostParams{
			ID:            uuid.New(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
			Title:         post.Title,
			Link:          post.Link,
			Description:   post.Description,
			PublishedDate: PubDate,
			FeedID:        DBFeed.ID,
		})
	}

	// update feed last_fetched_at field
	db.UpdateLastFetchedAt(context.Background(), DBFeed.ID)

}
