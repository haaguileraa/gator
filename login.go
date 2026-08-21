package main

import (
	"context"
	"fmt"
)

func handlerLogin(st *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("command usage: %s <username>", cmd.Name)
	}
	username := cmd.Args[0]
	
	user, err := st.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not find username %s, %w", username, err)
	}
	
	err = st.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set username %s, %w", user.Name, err)
	}

	fmt.Println("user set to", st.cfg.CurrentUserName)
	return nil
}
