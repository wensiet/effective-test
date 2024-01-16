package main

import (
	_ "effective-test/docs"
	"effective-test/internal/api"
	"effective-test/internal/config"
	"effective-test/internal/models"
)

func init() {
	err := config.Get().Database.AutoMigrate(&models.Person{})
	if err != nil {
		panic(err)
	}
}

// @title           			Effective-test
// @version         			1.0
// @description     			Test for effective company.

// @BasePath /api/v1
func main() {
	api.Run()
}
