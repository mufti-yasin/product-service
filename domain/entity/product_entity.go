package entity

import (
	valueobject "item-service/domain/value_object"
	"time"
)

// entity product
type Product struct {
	name        string
	category    []string
	quantity    float64
	price       float64
	imageUrl    string
	sku         valueobject.SKU
	description string
	createdAt   time.Time
	updatedAt   *time.Time
	deletedAt   *time.Time
}

type ProductDTO struct {
	Name        string
	Category    []string
	Quantity    float64
	Price       float64
	ImageUrl    string
	Sku         valueobject.SKU
	Description string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

// initial new product
func NewProduct(dto ProductDTO) Product {
	return Product{
		name:        dto.Name,
		category:    dto.Category,
		price:       dto.Price,
		imageUrl:    dto.ImageUrl,
		sku:         dto.Sku,
		quantity:    dto.Quantity,
		description: dto.Description,
		createdAt:   dto.CreatedAt,
		updatedAt:   dto.UpdatedAt,
		deletedAt:   dto.DeletedAt,
	}
}

// setter
func (e *Product) SetName(name string) {
	e.name = name
}

func (e *Product) SetCategory(categories ...string) {
	e.category = categories
}

func (e *Product) SetPrice(price float64) {
	e.price = price
}

func (e *Product) SetImageUrl(url string) {
	e.imageUrl = url
}

func (e *Product) SetSku(sku valueobject.SKU) {
	e.sku = sku
}

func (e *Product) SetQuantity(qty float64) {
	e.quantity = qty
}

func (e *Product) SetDescription(desc string) {
	e.description = desc
}

func (e *Product) SetCreatedAt(createdAt time.Time) {
	e.createdAt = createdAt
}

func (e *Product) SetUpdatedAt(updatedAt *time.Time) {
	e.updatedAt = updatedAt
}

func (e *Product) SetDeletedAt(deletedAt *time.Time) {
	e.deletedAt = deletedAt
}

// getter
func (e *Product) GetName() string {
	return e.name
}

func (e *Product) GetCategory() []string {
	return e.category
}

func (e *Product) GetPrice() float64 {
	return e.price
}

func (e *Product) GetImageUrl() string {
	return e.imageUrl
}

func (e *Product) GetSku() string {
	return e.sku.GetValue()
}

func (e *Product) GetQuantity() float64 {
	return e.quantity
}

func (e Product) GetDescription() string {
	return e.description
}

func (e Product) GetCreatedAt() time.Time {
	return e.createdAt
}

func (e Product) GetUpdatedAt() *time.Time {
	return e.updatedAt
}

func (e Product) GetDeletedAt() *time.Time {
	return e.deletedAt
}
