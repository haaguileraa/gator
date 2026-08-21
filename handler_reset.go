package main

import (
	"context"
	"fmt"
)

// Please do not do this in production. It is just for test purposes!

func handlerReset(st *state, cmd command) error {
	err := st.db.CleanDatabase(context.Background())
	if err != nil {
		return fmt.Errorf("could not reset database, %w", err)
	}

	fmt.Println("database reset succesfully")
	return nil
}


