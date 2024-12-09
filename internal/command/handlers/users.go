package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
)

func Users(s *internal.State, cmd command.Command) error {
	users, err := s.Queries.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error get the list of users Error: %s", err.Error())
	}

	for _, user := range users {
		name := user
		if user == s.Cfg.CurrentUserName {
			name += " (current)"
		}

		fmt.Println(name)
	}
	return nil
}

func ResetUsers(s *internal.State, cmd command.Command) error {
	_, err := s.Queries.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error resetting the users table Error: %s", err.Error())
	}

	fmt.Println("Reset all users in the users table")
	return nil
}

func Login(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("login expects a single parameter for the user name. command: %s", cmd)
	}

	user, err := s.Queries.GetUserByName(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error retrieving user %s. Error: %s", cmd.Args[0], err.Error())
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has logged in.\n", user.Name)
	return nil
}

func Register(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("register expects a single parameter for the user. command: %s", cmd)
	}

	uuid := uuid.New()
	now := time.Now()
	name := cmd.Args[0]

	user, err := s.Queries.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
	if err != nil {
		return err
	}

	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been registered.\n", user.Name)
	return nil
}
