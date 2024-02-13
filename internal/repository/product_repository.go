package repository

import (
	"context"
	"item-service/domain/entity"
	"item-service/domain/repository"
	"item-service/internal/repository/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type mongoProductRepo struct {
	db                *mongo.Database
	productCollection *mongo.Collection
}

func NewMongoBusinessRepo(db *mongo.Database) repository.ProductRepository {
	return &mongoProductRepo{
		db:                db,
		productCollection: db.Collection("product"),
	}
}

func (r *mongoProductRepo) Create(ctx context.Context, product entity.Product) error {

	var productModel = model.ProductModel{
		Name:        product.GetName(),
		Category:    product.GetCategory(),
		Quantity:    product.GetQuantity(),
		Price:       product.GetPrice(),
		ImageUrl:    product.GetImageUrl(),
		Sku:         product.GetSku(),
		Description: product.GetDescription(),
		CreatedAt:   product.GetCreatedAt(),
	}

	var _, err = r.productCollection.InsertOne(ctx, productModel)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}
	return nil
}
