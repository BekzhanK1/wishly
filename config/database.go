package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EnvDBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewEnvDBConfig() *EnvDBConfig {
	return &EnvDBConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_DATABASE"),
	}
}

func (c *EnvDBConfig) GetHost() string {
	return c.host
}

func (c *EnvDBConfig) GetPort() string {
	return c.port
}

func (c *EnvDBConfig) GetUsername() string {
	return c.username
}

func (c *EnvDBConfig) GetPassword() string {
	return c.password
}

func (c *EnvDBConfig) GetDatabase() string {
	return c.database
}

func ConnectDB(dbConfig EnvDBConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbConfig.GetUsername(), dbConfig.GetPassword(), dbConfig.GetPort(), dbConfig.GetPassword(), dbConfig.GetDatabase())
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
