package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
)

func handlerFollowing(st *state, cmd command) error {
	follows, err := st.db.GetFollowsForUser(context.Background(), st.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error obtaining follows for the current user: %w", err)
	}

	for _, follow := range follows {
		printFollow(follow)
	}	

	return nil
}

func printFollow(follow database.GetFollowsForUserRow) {
	fmt.Println("- ID:		", follow.ID)
	fmt.Println("- Created at:	", follow.CreatedAt)
	fmt.Println("- Updated at:	", follow.UpdatedAt)
	fmt.Println("- User ID:		", follow.UserID)
	fmt.Println("- User name:	", follow.UserName)
	fmt.Println("- FeedID:		", follow.FeedID)
	fmt.Println("- Feed name:	", follow.FeedName)
}
