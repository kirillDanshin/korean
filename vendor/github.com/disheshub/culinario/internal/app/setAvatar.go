package app

import (
	"io"
)

type SaveParams struct {
	File io.ReadCloser
	ID   int
}

func (app *app) SetAvatar(ctx Ctx, log Log, userID int, avatarFile io.ReadCloser) (user *User, err error) {
	return app.db.UpdateAvatar(ctx, log, userID, avatarFile)
}
