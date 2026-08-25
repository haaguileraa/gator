package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/haaguileraa/gator/internal/database"
	"time"
)

func handlerAddFeed(st *state, cmd command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("command usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

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

	ctx := context.Background()
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

