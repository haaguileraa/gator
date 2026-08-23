package main 

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

const appName = "gator"

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("user-agent", appName)

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error sending request: %w", err)
	}
		
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, fmt.Errorf("error reading response body: %w", err)
	}
	
	var rssFeed RSSFeed
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return &RSSFeed{}, fmt.Errorf("error unmarshaling data: %w", err)
	}
	rssFeed = unescapeRSSFeed(rssFeed)
	return &rssFeed, nil 
}


func unescapeRSSFeed(rssf RSSFeed) RSSFeed {
	rssf.Channel.Title = html.UnescapeString(rssf.Channel.Title)
	rssf.Channel.Description = html.UnescapeString(rssf.Channel.Description)
	for i := range len(rssf.Channel.Item) { 
		rssf.Channel.Item[i].Title = html.UnescapeString(rssf.Channel.Item[i].Title)
		rssf.Channel.Item[i].Description = html.UnescapeString(rssf.Channel.Item[i].Description)
	}
	return rssf
}
