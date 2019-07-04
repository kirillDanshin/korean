package app

type UpdateUserInfo struct {
	UserID   int
	Username string
}

func (app *app) UpdateUsername(ctx Ctx, log Log, userID int, newUsername string) (user *User, err error) {
	return app.db.UpdateUsername(ctx, log, userID, newUsername)
}
