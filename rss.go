package main

import (
	"encoding/xml"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string `xml:"name"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Items       []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
}

func URLToRSS(url string) (RSSFeed, error) {
	client := http.Client{}
	res, err := client.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}

	data_byte, err := io.ReadAll(res.Body)
	if err != nil {
		return RSSFeed{}, err
	}

	Feed := RSSFeed{}
	xml.Unmarshal(data_byte, &Feed)
	return Feed, nil
}
