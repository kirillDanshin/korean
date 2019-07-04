package app

import (
	"github.com/powerman/structlog"
	"golang.org/x/net/context"
	"io"
)

type (
	// DB - main repository containing user and recipe information
	DB interface {
		ExistsEmail(ctx Ctx, log Log, usersEmail string) (exists bool, err error)
		CreateUser(ctx Ctx, log Log, user NewUser) (ID int, err error)
		GetPassAndIDByEmail(ctx Ctx, log Log, userEmail string) (ID int, passHash string, err error)

		UpdateUsername(ctx Ctx, log Log, userID int, newUsername string) (user *User, err error)
		UpdateAvatar(ctx Ctx, log Log, ID int, file io.ReadCloser) (user *User, err error)
		GetUserByID(ctx Ctx, log Log, userID int) (user *User, err error)

		CreateMeasureWeight(ctx Ctx, log Log, measureWeight NewMeasureWeight) (listMeasureWeight []MeasureWeight, err error)
		RemoveMeasureWeight(ctx Ctx, log Log, ID int) error

		CreateIngredient(ctx Ctx, log Log, ingredient NewIngredient) (listIngredient []Ingredient, err error)
		GetIngredient(ctx Ctx, log Log) (listIngredient []Ingredient, err error)
		RemoveIngredient(ctx Ctx, log Log, ID int) error

		CreateRecipe(ctx Ctx, log Log, recipeInfo NewRecipe) (*Recipe, error)
		GetRecipe(ctx Ctx, log Log, recipeID int) (*Recipe, error)
	}
	// SessionStorage - containing user sessions
	SessionStorage interface {
		Set(token string, id int) (err error)
		GetID(token string) (id int, err error)
	}
	// App - provides application features service.
	App interface {
		Registration(ctx Ctx, log Log, newUser NewUser) (token string, err error)
		Login(ctx Ctx, log Log, loginInfo LoginInfo) (token string, err error)
		CheckEmail(ctx Ctx, log Log, email string) (exists bool, err error)

		CheckAuth(token string) (*Session, error)

		GetUser(ctx Ctx, log Log, userID int) (user *User, err error)
		UpdateUsername(ctx Ctx, log Log, userID int, newUsername string) (user *User, err error)
		SetAvatar(ctx Ctx, log Log, userID int, avatarFile io.ReadCloser) (user *User, err error)

		CreateMeasureWeight(ctx Ctx, log Log, measureWeight NewMeasureWeight) (listMeasureWeight []MeasureWeight, err error)
		RemoveMeasureWeight(ctx Ctx, log Log, ID int) error

		CreateIngredient(ctx Ctx, log Log, ingredient NewIngredient) (listIngredient []Ingredient, err error)
		GetIngredient(ctx Ctx, log Log) (listIngredient []Ingredient, err error)
		RemoveIngredient(ctx Ctx, log Log, ID int) error

		CreateRecipe(ctx Ctx, log Log, recipeInfo NewRecipe) (*Recipe, error)
		GetRecipe(ctx Ctx, log Log, recipeID int) (*Recipe, error)
	}

	app struct {
		db      DB
		session SessionStorage
	}
	// Сtx is a synonym for convenience.
	Ctx = context.Context
	// Log is a synonym for convenience.
	Log = *structlog.Logger
)

// New - return new application.
func New(db DB, ss SessionStorage) App {
	return &app{db: db, session: ss}
}
