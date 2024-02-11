package config

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		DB
		JWT
		Mongo
		GRPC `yaml:"grpc"`
		Cloudinary
	}

	// App -.
	App struct {
		Name     string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version  string `env-required:"true" yaml:"version" env:"APP_VERSION"`
		TimeZone string `yaml:"timezone" env:"APP_TIMEZONE" `
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// JWT -.
	JWT struct {
		SECRET string `env-required:"true" env:"JWT_SECRET"`
		TTL    int64  `yaml:"jwt_ttl" env:"JWT_TTL" env-default:"60"`
	}

	// DB
	DB struct {
		URL string `env-required:"true" env:"DB_URL"`
	}

	// Mongo
	Mongo struct {
		DSN string `env-required:"true" env:"MONGO_DSN"`
		DB  string `env-required:"true" env:"MONGO_DB"`
	}

	// gRPC
	GRPC struct {
		Port string `env:"GRPC_PORT" yaml:"port"`
	}

	Cloudinary struct {
		CloudName string `env:"CLOUDINARY_CLOUD_NAME"`
		ApiKey    string `env:"CLOUDINARY_API_KEY"`
		ApiScret  string `env:"CLOUDINARY_API_SECRET"`
		Folder    string `env:"CLOUDINARY_UPLOAD_FOLDER"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Config error : ", err)
	}
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
