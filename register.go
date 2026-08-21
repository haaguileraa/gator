package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/haaguileraa/gator/internal/database"
	"time"
)

func handlerRegister(st *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("command usage: %s <username>", cmd.Name)
	}
	username := cmd.Args[0]

	createdAt := time.Now().UTC()

	params := database.CreateUserParams{
		ID: 		uuid.New(),
		CreatedAt:	createdAt,	
		UpdatedAt:	createdAt,
		Name:		username,
	}

	user, err := st.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("could not create username %s, %w", username, err)
	}
	
	err = st.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("user set to", st.cfg.CurrentUserName)
	return nil
}
