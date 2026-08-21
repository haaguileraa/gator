package main

import (
	"context"
	"fmt"
)

func handlerListUsers(st *state, cmd command) error {
	usernames, err := st.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not retrieve users, %w", err)
	}

	for _, username := range usernames {
		currentStr := ""
		if (username == st.cfg.CurrentUserName) {
			currentStr = "(current)"
		}
		fmt.Println("	*", username, currentStr)
	}

	return nil
}
