package repository

import (
	"context"
	"item-service/domain/entity"
)

type ProductRepository interface {
	//create data with product entity as parameter and return error
	Create(context.Context, entity.Product) error
}
