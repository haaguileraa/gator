package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/haaguileraa/gator/internal/database"
	"log"
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
		log.Println("error getting next feed:", err)
	}

	rssf, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		log.Printf("error fetching url %s: %w\n", feed.Url, err)
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
		log.Println("error marking feed last fetched:", err)
	}


	fmt.Println("- Title:", rssf.Channel.Title)
	for _, item := range rssf.Channel.Item {
		addPost(st, item, feed.ID)
	}
}


func addPost(st *state, item RSSItem, feedID uuid.UUID) {
	// description
	description := sql.NullString {
		String:	item.Description,
		Valid:	item.Description != "",
	}
	// pubDate
	t, err := time.Parse(time.RFC3339, item.PubDate)
	publishedAt := sql.NullTime {
		Time:	t,
		Valid:	err == nil,
	}
	
	entryTime := time.Now().UTC()
	params := database.CreatePostParams {
		ID: uuid.New(),
		CreatedAt:	entryTime,
		UpdatedAt:	entryTime,
		Title:		item.Title,
		Url:		item.Link,
		Description:	description,
		PublishedAt:	publishedAt,
		FeedID:		feedID,
	}
	
	post, err := st.db.CreatePost(context.Background(), params)
	if err != nil {
		logSqlError(err)
		return
	}

	log.Println("	- Title:", post.Title)
}
