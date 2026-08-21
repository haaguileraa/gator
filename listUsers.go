package main

import (
	"context"
	"fmt"
)

func handlerListUsers(st *state, cmd command) error {
	users, err := st.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not retrieve users, %w", err)
	}

	for _, user := range users {
		if (user.Name == st.cfg.CurrentUserName) {
			fmt.Printf("	* %s (current)", user.Name)
		}
		fmt.Println("	*", user.Name)
	}

	return nil
}
