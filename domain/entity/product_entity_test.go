package entity_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"item-service/domain/entity"
	valueobject "item-service/domain/value_object"
)

func TestProductSetName(t *testing.T) {
	product := entity.Product{}
	product.SetName("New Name")

	assert.Equal(t, "New Name", product.GetName(), "Name should be updated")
}

func TestProductSetCategory(t *testing.T) {
	product := entity.Product{}
	product.SetCategory("New Category")

	expectedCategories := []string{"New Category"}
	assert.Equal(t, expectedCategories, product.GetCategory(), "Categories should be updated")
}

func TestProductSetPrice(t *testing.T) {
	product := entity.Product{}
	product.SetPrice(15.0)

	assert.Equal(t, 15.0, product.GetPrice(), "Price should be updated")
}

func TestProductSetImageUrl(t *testing.T) {
	product := entity.Product{}
	product.SetImageUrl("new_image.jpg")

	assert.Equal(t, "new_image.jpg", product.GetImageUrl(), "Image URL should be updated")
}

func TestProductSetSku(t *testing.T) {
	sku := valueobject.NewSKU("sku", time.Now())
	product := entity.Product{}
	product.SetSku(sku)

	assert.Equal(t, sku, product.GetSku(), "SKU should be updated")
}

func TestProductSetQuantity(t *testing.T) {
	product := entity.Product{}
	product.SetQuantity(10.0)

	assert.Equal(t, 10.0, product.GetQuantity(), "Quantity should be updated")
}

func TestProductSetDescription(t *testing.T) {
	product := entity.Product{}
	product.SetDescription("New Description")

	assert.Equal(t, "New Description", product.GetDescription(), "Description should be updated")
}

func TestProductSetCreatedAt(t *testing.T) {
	createdAt := time.Now()
	product := entity.Product{}
	product.SetCreatedAt(createdAt)

	assert.Equal(t, createdAt, product.GetCreatedAt(), "Created at time should be updated")
}

func TestProductSetUpdatedAt(t *testing.T) {
	updatedAt := time.Now()
	product := entity.Product{}
	product.SetUpdatedAt(&updatedAt)

	assert.Equal(t, updatedAt, *product.GetUpdatedAt(), "Updated at time should be updated")
}

func TestProductSetDeletedAt(t *testing.T) {
	deletedAt := time.Now()
	product := entity.Product{}
	product.SetDeletedAt(&deletedAt)

	assert.Equal(t, deletedAt, *product.GetDeletedAt(), "Deleted at time should be updated")
}
