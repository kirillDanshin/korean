package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ZergsLaw/korean/internal/api"
	"github.com/ZergsLaw/korean/internal/db"
	"github.com/ZergsLaw/korean/internal/session"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/powerman/structlog"
)

// nolint:gochecknoglobals
var (
	log = structlog.New()

	cfg struct {
		api            api.Configuration
		db             db.Configuration
		sessionStorage session.Configuration
	}
)

func flagParse() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	flag.IntVar(&cfg.api.Port, "port", 8080, "Server port")
	flag.StringVar(&cfg.api.Host, "host", hostname, "Server host")

	// TODO remove, temporary solution
	flag.StringVar(&cfg.api.AdminUsername, "username", "value", "Admin username.")
	flag.StringVar(&cfg.api.AdminPass, "password", "252525", "Admin password.")

	flag.StringVar(&cfg.db.Host, "db.host", "localhost", "DB host")
	flag.IntVar(&cfg.db.Port, "db.port", 5432, "DB port")
	flag.StringVar(&cfg.db.Name, "db.name", "postgresql", "DB name")
	flag.StringVar(&cfg.db.User, "db.user", "postgresql", "DB user")
	flag.StringVar(&cfg.db.Password, "db.pass", "postgresql", "DB password")
	flag.StringVar(&cfg.db.Type, "db.type", "postgres", "DB type")

	flag.StringVar(&cfg.sessionStorage.Host, "sessionStore.host", "localhost", "DB session host")
	flag.StringVar(&cfg.sessionStorage.Port, "sessionStore.port", "6379", "DB session port")
	flag.StringVar(&cfg.sessionStorage.Password, "sessionStore.pass", "", "DB session password")
	flag.DurationVar(&cfg.sessionStorage.Expire, "sessionStore.expire", session.Default, "lifetime session")

	flag.Parse()
}

func main() {
	flagParse()
	log.Fatal(run())
}

func run() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.db.Host,
		cfg.db.Port,
		cfg.db.User,
		cfg.db.Password,
		cfg.db.Name,
	)

	conn, err := sqlx.Connect(cfg.db.Type, dsn)
	if err != nil {
		return err
	}

	store, err := db.New(conn)
	if err != nil {
		return err
	}

	sessionStore, err := session.New(cfg.sessionStorage)
	if err != nil {
		return err
	}

	server, err := api.Serve(log, store, sessionStore, cfg.api)
	if err != nil {
		return err
	}

	return errors.Wrapf(server.Serve(), "failed to web server")
}
