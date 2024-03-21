package app

import (
	"lambda-function/api"
	"lambda-function/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// we actually inialize our DB store
	// gets passed DOWN into the api handler

	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
