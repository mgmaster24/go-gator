package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/database"
	"github.com/mgmaster24/gator/internal/rss"
)

func AddFeed(s *internal.State, cmd command.Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Add expects two arguments the feed name and url. command: %s", cmd)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]
	user, err := s.Queries.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.Queries.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("Error creating feed from %s. Error: %s", url, err.Error())
	}

	fmt.Printf("Create feed %s at %s\n", name, url)
	fmt.Println(feed)
	return nil
}

func Aggregate(s *internal.State, cmd command.Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error fetching feed from %s. Error: %s", url, err.Error())
	}

	fmt.Println("Fetched feed from ", url)
	fmt.Println(feed)
	return nil
}

func Feeds(s *internal.State, cmd command.Command) error {
	feeds, err := s.Queries.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		username, err := s.Queries.GetUserNameById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Println("Feed Name:", feed.Name)
		fmt.Println("Feed URL:", feed.Url)
		fmt.Println("Feed Add By:", username)
	}

	return nil
}
