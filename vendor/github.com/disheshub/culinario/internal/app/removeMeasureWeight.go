package app

func (app *app) RemoveMeasureWeight(ctx Ctx, log Log, ID int) error {
	return app.db.RemoveMeasureWeight(ctx, log, ID)
}
