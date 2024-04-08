package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/esiebomaj/rssarg/internal/database"
	"github.com/google/uuid"
)

func ScrapeFeeds(db *database.Queries, DurationInMins int16, concurency int16) {
	duration := time.Duration(DurationInMins) * time.Minute
	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
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
		return
	}

	// save posts in db
	for _, post := range rssFeed.Channel.Items {
		layout := "Mon, 02 Jan 2006 15:04:05 -0700"
		PubDate, err := time.Parse(layout, post.PubDate)
		if err != nil {
			fmt.Printf("Cannot parse time %v for post %v", post.PubDate, post.Title)
			continue
		}
		DBpost, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:            uuid.New(),
			CreatedAt:     time.Now().UTC(),
			UpdatedAt:     time.Now().UTC(),
			Title:         post.Title,
			Link:          post.Link,
			Description:   post.Description,
			PublishedDate: PubDate,
			FeedID:        DBFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "posts_link_unique") {
				continue
			}
			fmt.Printf("could not save <%v> to DB: %v \n", post.Title, err)
			continue
		}
		fmt.Println("New post discovered:", DBpost.Title)
	}
	fmt.Printf("\n*************** discovered %v posts for '%v' \n", len(rssFeed.Channel.Items), rssFeed.Channel.Title)

	// update feed last_fetched_at field
	db.UpdateLastFetchedAt(context.Background(), DBFeed.ID)

}
