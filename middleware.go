package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
)

func middlewareLoggedIn(handler func(st *state, cmd command, user database.User) error) func(*state, command) error {	
	return func(st *state, cmd command) error {
		user, err := st.db.GetUser(context.Background(), st.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("could not find username %s, %w", user.Name, err)
		}
		return handler(st, cmd, user)
	}
}
