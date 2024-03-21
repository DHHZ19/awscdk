package api

import (
	"encoding/json"
	"fmt"
	"lambda-function/database"
	"lambda-function/types"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ApiHandler struct {
	dbStore database.UserStore
}

func NewApiHandler(dbStore database.UserStore) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var registerUser types.RegisterUser

	err := json.Unmarshal([]byte(request.Body), &registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	if registerUser.Username == "" || registerUser.Passsword == "" {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid request - fields empty",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	// does a user with this username already exists?
	userExists, err := api.dbStore.DoesUserExist(registerUser.Username)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server errror",
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	if userExists {
		return events.APIGatewayProxyResponse{
			Body:       "User already existsty",
			StatusCode: http.StatusConflict,
		}, err
	}

	user, err := types.NewUser(registerUser)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("errror instering user - %w", err)
	}

	// we know that a user does not exist
	err = api.dbStore.InsertUser(user)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "internal server error",
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("errror instering user - %w", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       "succesfuly regsitered user",
		StatusCode: http.StatusOK,
	}, nil
}

func (api ApiHandler) LoginUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var loginRequest LoginRequest

	err := json.Unmarshal([]byte(request.Body), &loginRequest)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Request",
			StatusCode: http.StatusBadRequest,
		}, err
	}

	user, err := api.dbStore.GetUser(loginRequest.Username)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Internal server error",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	if !types.ValidatePassword(user.PasswrodHash, loginRequest.Password) {
		return events.APIGatewayProxyResponse{
			Body:       "invalid user credentials",
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Sucessfuly logged in",
		StatusCode: http.StatusOK,
	}, nil
}
