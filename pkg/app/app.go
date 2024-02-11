package app

import (
	"item-service/config"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// App DB
type Database struct {
	Gorm  *gorm.DB
	Mongo *mongo.Database
}

// App
type App struct {
	DB     Database
	Config *config.Config
}
