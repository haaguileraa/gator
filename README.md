# gator
Building a Blog Aggregator (Gator) in Go

# Prerequisites

1. To run this CLI tool you need the [Go Toolchain](https://go.dev/doc/install) with version 1.25 or higher.
2. You need to install all the requirements to work with [PostgreSQL](https://www.postgresql.org/). Here is a very useful [documentation](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql) for WSL.
3. After setting up your database, you can run ```$ go install``` at the root of the project.
4. You will need a configuration file in your home directory named ```~/.gatorconfig.json```. The content should look like this: 
``` 
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
5. You should be able to run the program using ```$ gator <command> <argument>``` 

# Commands

- ```addFeed <name> <url>```: adds a new feed.
- ```agg <time between requests>```: aggregates posts from the feed URLs.
- ```browse [limit]```: prints the posts from the current logged-in user with a limit given by the second argument. It defaults to 2 if left empty.
- ```feeds```: prints out all feeds and the user following them.
- ```feedscurrent```: prints out all feeds followed by the current user.
- ```follow <url>```: adds a new url to follow for the current user.
- ```following```: prints out all feeds followed by the current user.
- ```login <username>```: logs in the given user.
- ```register <username>```: adds (registers) a new user to the database.
- ```reset```: *DANGER!* this command... yes, resets the database without asking for confirmation :)
- ```unfollow <url>```: lets the current user to unfollow a feed by url. 
- ```users```: prints out all registered users.


Have fun!
