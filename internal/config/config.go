package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

type config struct {
	Database *gorm.DB
}

var conf config
var once sync.Once

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASS"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PORT"),
	)
	conf.Database, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
}

func Get() config {
	once.Do(initConfig)
	return conf
}
