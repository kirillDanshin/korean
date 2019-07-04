package app

func (app *app) GetUser(ctx Ctx, log Log, userID int) (user *User, err error) {
	return app.db.GetUserByID(ctx, log, userID)
}
