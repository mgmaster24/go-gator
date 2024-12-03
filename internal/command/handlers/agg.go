package handlers

import (
	"context"
	"fmt"

	"github.com/mgmaster24/gator/internal"
	"github.com/mgmaster24/gator/internal/command"
	"github.com/mgmaster24/gator/internal/rss"
)

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
