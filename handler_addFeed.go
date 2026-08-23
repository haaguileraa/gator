package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
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

	params := database.CreateFeedParams{
		ID: user.ID,
		Name: name,
		Url: url,
	}
	feeds, err := st.db.CreateFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	
	fmt.Printf("feed created successfully\n%+v", feeds)
	return nil
}
