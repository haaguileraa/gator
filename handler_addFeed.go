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
	
	createdAt := time.Now().UTC()

	params := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
		Name: name,
		Url: url,
		UserID: user.ID,
	}

	feeds, err := st.db.CreateFeed(ctx, params)
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}
	
	fmt.Printf("feed created successfully\n%+v", feeds)
	return nil
}
