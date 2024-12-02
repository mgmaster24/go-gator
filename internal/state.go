package internal

import (
	"github.com/mgmaster24/gator/internal/config"
	"github.com/mgmaster24/gator/internal/database"
)

type State struct {
	Cfg     *config.Config
	Queries *database.Queries
}
