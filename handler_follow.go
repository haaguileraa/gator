package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/haaguileraa/gator/internal/database"
	"time"
)

func handlerFollow(st *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command usage: %s <url>", cmd.Name)
	}
	
	ctx := context.Background()

	url := cmd.Args[0]
	feed, err := st.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("error getting feed from url '%s': %w", url, err)
	}
	
	createdAt := time.Now().UTC()

	params := database.CreateFeedFollowParams{
		ID:		uuid.New(),
		CreatedAt:	createdAt,
		UpdatedAt:	createdAt,
		UserID:		user.ID,
		FeedID:		feed.ID,
	}

	inserted, err := st.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating follow record for the current user: %w", err)
	}

	fmt.Printf("user %s is now following feed %s", inserted.UserName, inserted.FeedName)	
	return nil
}
