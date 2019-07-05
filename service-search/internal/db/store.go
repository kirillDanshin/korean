package db

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/powerman/structlog"
)

type (
	storage struct {
		elastic *elastic.Client
	}

	// Configuration - contains configuration for storage.
	Configuration struct {
		URL string
	}

	// Ctx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)


func (s *storage) GetByID(ctx Ctx, id int) {

}
