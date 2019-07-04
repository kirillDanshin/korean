package app

type NewMeasureWeight struct {
	Name string
}

func (app *app) CreateMeasureWeight(ctx Ctx, log Log, measureWeight NewMeasureWeight) (listMeasureWeight []MeasureWeight, err error) {
	return app.db.CreateMeasureWeight(ctx, log, measureWeight)
}
