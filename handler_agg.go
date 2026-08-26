package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
	"time"
)

func handlerAgg(st *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command usage: %s <time between requests>\nexample: agg 1m", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("collecting feeds every", timeBetweenReqs)
	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(st, user)
	} 
	return nil
}

func scrapeFeeds(st *state, user database.User) {
	ctx := context.Background()
	feed, err := st.db.GetNextFeedToFetch(ctx, user.ID)
	if err != nil {
		fmt.Println("error getting next feed:", err)
	}

	rssf, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		fmt.Printf("error fetching url %s: %w\n", feed.Url, err)
	}

	fmt.Println("- Title:", rssf.Channel.Title)
	for _, item := range rssf.Channel.Item {
		fmt.Println("	- Title:", item.Title)
	}


	fetchedTime := time.Now().UTC() 
	
	lastFetchedAt := sql.NullTime{
		Time: 	fetchedTime,
		Valid: 	true,
	}

	markParams := database.MarkFeedFetchedParams{
		ID: 		feed.ID,
		UpdatedAt: 	fetchedTime,
		LastFetchedAt:	lastFetchedAt,
	}

	err = st.db.MarkFeedFetched(ctx, markParams)
	if err != nil {
		fmt.Println("error marking feed last fetched:", err)
	}
}


