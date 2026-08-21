package main

import(
	"fmt"
	"github.com/haaguileraa/gator/internal/config"
	"github.com/haaguileraa/gator/internal/database"
)

type state struct {
	db	*database.Queries
	cfg	*config.Config
}

func (s *state) SetUser(user string) error {
	err := s.cfg.SetUser(user)
	if err != nil {
		return fmt.Errorf("could not set user: %v", err)
	}

	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("could not read config back: %v", err)
	}

	s.cfg = &cfg
	return nil
}
