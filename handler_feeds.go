package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
)

func handlerFeeds(st *state, cmd command) error {
	ctx := context.Background()

	feeds, err := st.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error obtaining feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds to show")
		return nil
	}
	
	for _, feed := range feeds {
		user, err := st.db.GetUserByID(ctx, feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user ID '%d' information: %w", feed.UserID, err)
		}
		fmt.Println("user name:", user.Name)
		printFeed(feed)
	}
	return nil
}

func handlerFeedsCurrentUser(st *state, cmd command, user database.User) error {	
	feeds, err := st.db.GetFeedsByUserId(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error obtaining feeds for user %s: %w", user.Name, err)
	}
	
	if len(feeds) == 0 {
		fmt.Println("no feeds to show")
		return nil
	}

	fmt.Println("feeds from user", user.Name)
	for _, feed := range feeds {
		printFeed(feed)
	}
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Println("- ID:		", feed.ID)
	fmt.Println("- Created at:	", feed.CreatedAt)
	fmt.Println("- Updated at:	", feed.UpdatedAt)
	fmt.Println("- Name:		", feed.Name)
	fmt.Println("- URL:		", feed.Url)
	fmt.Println("- User ID:		", feed.UserID)
}
