package main

import (
	"context"
	"fmt"
	"github.com/haaguileraa/gator/internal/database"
	"strconv"
)

func handlerBrowse(st *state, cmd command, user database.User) error {
	var limit int32

	if len(cmd.Args) != 1 {
		limit = 2
	} else {
		parsed, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("cannot parse %s as limit value", cmd.Args[0])
		}
		limit = int32(parsed)
	}

	params := database.GetPostsForUserParams {
		UserID:	user.ID,
		Limit:	limit,
	}

	posts, err := st.db.GetPostsForUser(context.Background(), params)

	if err != nil {
		return fmt.Errorf("error obtaining posts for user %s: %w", user.Name, err)
	}

	for _, post := range posts {
		printPost(post)
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Println("- Title:", post.Title)
	if post.PublishedAt.Valid {
		fmt.Println("- Published at:", post.PublishedAt.Time)
	}
	if post.Description.Valid {
		fmt.Println("- Description:", post.Description.String)
	}
	fmt.Println()
}
