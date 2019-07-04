package app

type NewIngredient struct {
	Name            string
	MeasureWeightID int
}

func (app *app) CreateIngredient(ctx Ctx, log Log, ingredient NewIngredient) (listIngredient []Ingredient, err error) {
	return app.db.CreateIngredient(ctx, log, ingredient)
}
