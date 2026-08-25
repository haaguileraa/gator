package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
)

func handlerUnfollow(st *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("command usage: %s <url>", cmd.Name)
	}

	ctx := context.Background()
	url := cmd.Args[0]
	feed, err := st.db.GetFeedByUrl(ctx, url)
	if err != nil {
		return fmt.Errorf("error getting feed from url '%s': %w", url, err)
	}

	params := database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	err = st.db.Unfollow(ctx, params) 
	if err != nil {
		return fmt.Errorf("error unfollowing feed %s with url %s for user %s: %w", feed.Name, feed.Url, user.Name, err)
	}
	fmt.Println("feed %s with url %s deleted successfully for user %s")
	return nil
}
