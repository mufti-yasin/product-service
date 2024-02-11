package usecase

import (
	"item-service/pkg/app"
)

type UseCase struct {
}

func New(app *app.App) *UseCase {
	return &UseCase{}
}
