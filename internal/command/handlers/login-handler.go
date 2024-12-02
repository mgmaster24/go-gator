package handlers

import (
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
)

func Login(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expexts at least a single parameter. command: %s", cmd)
	}

	err := s.Cfg.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", s.Cfg.CurrentUserName)
	return nil
}
