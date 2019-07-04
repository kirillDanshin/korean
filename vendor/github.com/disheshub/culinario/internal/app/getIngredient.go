package app

func (app *app) GetIngredient(ctx Ctx, log Log) (listIngredient []Ingredient, err error) {
	return app.db.GetIngredient(ctx, log)
}
