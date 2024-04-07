package main

import (
	"context"
	"fmt"
	"time"

	"github.com/esiebomaj/rssarg/internal/database"
)

func ScrapeFeeds(db *database.Queries, DurationInMins int16, concurency int16) {
	duration := time.Duration(DurationInMins) * time.Second
	ticker := time.NewTicker(duration)

	for range ticker.C {
		FetchFeeds(db, concurency)
	}

}

func FetchFeeds(db *database.Queries, limit int16) {
	// feed, err := URLToRSS("https://blog.boot.dev/index.xml")

	// fetch limit feeds for db
	fmt.Println("Fetching Feeds from db")
	feeds, err := db.GetEarliestFetchedFeeds(context.Background(), int32(limit))
	if err != nil {
		fmt.Println("err fetching feeds from db", err)
		return
	}
	fmt.Println(feeds)

	// fetch post for this feeds and save them in db

	// update feeds last_fetched_at field
}
