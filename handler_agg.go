package main

import (
	"context"
	"fmt"
)

func handlerAgg(st *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	rssf, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching url %s: %w", feedURL, err)
	}
	fmt.Printf("RSS Feed obtained successfully\n%+v", rssf)
	return nil
}
