package main

import(
	"fmt"
	"github.com/haaguileraa/gator/internal/config"
	"log"
)



func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	fmt.Printf("Config: %+v\n", cfg)

	err = cfg.SetUser("haaguileraa")
	if err != nil {
		log.Fatalf("could not set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}

	fmt.Printf("Config: %+v\n",cfg)
}
