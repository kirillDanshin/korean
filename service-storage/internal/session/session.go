package session

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"net"
	"strconv"
	"time"
)

const (
	// Day - one day
	Day = time.Hour * 24
	// Month - one month
	Month = Day * 30
	// Default expire session
	Default = Month
)

type (
	// Configuration - configuration for connection database
	Configuration struct {
		Host     string
		Port     string
		Password string
		Expire   time.Duration
	}

	Store interface {
		Set(ctx Ctx, token string, id int) (err error)
		GetID(token string) (id *int, err error)
	}

	sessionStorage struct {
		expire  time.Duration
		storage *redis.Client
	}

	// Сtx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

var (
	UserNotAuthorized = errors.New("User unauthorized")
)

// Option - building params for connection Redis
func (cfg *Configuration) Option() *redis.Options {
	return &redis.Options{
		Addr:     net.JoinHostPort(cfg.Host, cfg.Port),
		Password: cfg.Password,
	}
}

// New - building client for connection DB
func New(cfg Configuration) (*sessionStorage, error) {
	redisClient := redis.NewClient(cfg.Option())
	if err := redisClient.Ping().Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to connection redis")
	}

	return &sessionStorage{storage: redisClient, expire: cfg.Expire}, nil
}

// Set - new session in DB
func (c *sessionStorage) Set(ctx Ctx, token string, id int) (err error) {
	err = c.storage.WithContext(ctx).Set(token, id, c.expire).Err()

	return errors.Wrapf(err, "failed to save %s by key %d ", token, id)
}

// GetID - get id from DB
func (c *sessionStorage) GetID(token string) (*int, error) {
	value, err := c.storage.Get(token).Result()
	if err == redis.Nil {
		return nil, UserNotAuthorized
	}

	id, err := strconv.Atoi(value)
	if err != nil {
		return nil, errors.Wrapf(err, "id not int")
	}

	return &id, nil
}
