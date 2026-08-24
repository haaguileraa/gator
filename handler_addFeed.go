package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/haaguileraa/gator/internal/database"
	"time"
)

func handlerAddFeed(st *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("command usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	ctx := context.Background()
	
	user, err := st.db.GetUser(ctx, st.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting current user '%s' information: %w", st.cfg.CurrentUserName, err)
	}
	
	// create feed
	createdAt := time.Now().UTC()

	paramsFeed := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Name: name,
		Url: url,
		UserID: user.ID,
	}

	feed, err := st.db.CreateFeed(ctx, paramsFeed)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	
	fmt.Println("feed created successfully")
	printFeed(feed)

	// follow feed
	createdAt = time.Now().UTC()

	paramsFeedFollow := database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt:	createdAt,
		UpdatedAt:	createdAt,
		UserID:		user.ID,
		FeedID:		feed.ID,
	}

	inserted, err := st.db.CreateFeedFollow(ctx, paramsFeedFollow)
	if err != nil {
		return fmt.Errorf("error creating follow record for the current user: %w", err)
	}

	fmt.Printf("user %s is now following feed %s", inserted.UserName, inserted.FeedName)	
	return nil
}

