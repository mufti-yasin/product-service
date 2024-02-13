package repository_test

import (
	"context"
	"fmt"
	"item-service/domain/entity"
	valueobject "item-service/domain/value_object"
	"item-service/internal/repository"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCreateProduct(t *testing.T) {

	testCases := []struct {
		name     string
		entity   entity.Product
		expected any
		isError  bool
	}{
		{
			name: "positive case: when success create product on mongo",
			entity: entity.NewProduct(entity.ProductDTO{
				Name:        "headset",
				Category:    []string{"elektronik", "gadget"},
				Quantity:    10,
				Price:       50200,
				ImageUrl:    "www.gogle.com",
				Sku:         valueobject.NewSKU("test_product", time.Now()),
				Description: "made in cina",
				CreatedAt:   time.Now(),
			}),
			expected: nil,
			isError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			var (
				ctx = context.TODO()

				db = ConnectMongo(ctx)

				productEntity = testCase.entity

				repository = repository.NewMongoBusinessRepo(db)

				err = repository.Create(ctx, productEntity)
			)

			if testCase.isError {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

func ConnectMongo(ctx context.Context) *mongo.Database {
	// Connect
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://yalqurni2:7nrMf8dMyr1me2qt@product-service.9esrnyp.mongodb.net"))
	if err != nil {
		log.Fatal("Mongo database connection error ", err)
	}

	// Set database
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo database connection error ", err)
	}
	fmt.Println("Mongo database connection successfully")
	db := client.Database("testing")
	return db
}
