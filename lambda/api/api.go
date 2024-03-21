package api

import (
	"fmt"
	"lambda-function/database"
	"lambda-function/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Passsword == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// does a user with this username already exists?
	userExists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("there an error checking if user exists %w", err)
	}

	if userExists {
		return fmt.Errorf("a user wit that username already exissts")
	}

	// we know that a user does not exist
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error regsitering the user %w", err)
	}

	return nil
}
