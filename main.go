package main

import(
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/haaguileraa/gator/internal/config"
	"github.com/haaguileraa/gator/internal/database"
	"log"
	"os"
	"fmt"
)


func main() {
	// load config
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	fmt.Println("connecting to database", cfg.DbURL)

	// open the connection to the database
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)

	st := state{
		cfg :	&cfg,
		db :	dbQueries,
	}

	if err != nil {
		log.Fatalf("Error creating new state %v", err)
	}
	cmds := NewCommands()
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("feedscurrent", middlewareLoggedIn(handlerFeedsCurrentUser))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Expecting at least 2 arguments, got %d instead", len(args))
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}
	err = cmds.run(&st, cmd)
	if err != nil {
		log.Fatalf("Could not execute command %+v\n%v", cmd, err)
	}
}
