package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlxPatidar/go-restful-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Collection
}

func NewUserController(client *mongo.Collection) *UserController {
	return &UserController{client}
}

// get all users
func (userController UserController) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all user")
	res.Header().Add("Content-Type", "application/json")
	users := models.Users(userController.client)
	json.NewEncoder(res).Encode(users)
}

// get user information based on user id
func (userController UserController) GetUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get user")
	res.Header().Add("Content-Type", "application/json")
	params := mux.Vars(req)
	userId := params["userId"]
	if !primitive.IsValidObjectID(userId) {
		json.NewEncoder(res).Encode("User id not found")
	}
	_id, _ := primitive.ObjectIDFromHex(userId)
	user := models.UserOne(userController.client, _id)
	if user == nil {
		json.NewEncoder(res).Encode("No user found")
		return
	}
	json.NewEncoder(res).Encode(user)
}

// create new user
func (userController UserController) CreateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create user")
	res.Header().Add("Content-Type", "application/json")
	if req.Body == nil {
		json.NewEncoder(res).Encode("Please send some data.")
	}
	var newUser models.User
	_ = json.NewDecoder(req.Body).Decode(&newUser)
	user := models.CreateUser(userController.client, newUser)
	if user == false {
		json.NewEncoder(res).Encode("Invalid user data provided.")
	}
	json.NewEncoder(res).Encode("New user created successfully.")
}

// update user detail
func (userController UserController) UpdateUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Update user")
	res.Header().Add("Content-Type", "application/json")
	userId := mux.Vars(req)["userId"]
	if !primitive.IsValidObjectID(userId) {
		json.NewEncoder(res).Encode("User id not found")
	}
	_id, _ := primitive.ObjectIDFromHex(userId)
	user := models.UpdateUser(userController.client, _id, primitive.M{})
	if user == false {
		json.NewEncoder(res).Encode("Invalid user data provided.")
	}
	json.NewEncoder(res).Encode("User information updated successfully.")
}

// update user detail
func (userController UserController) DeleteUser(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Delete user")
	res.Header().Add("Content-Type", "application/json")
	userId := mux.Vars(req)["userId"]
	if !primitive.IsValidObjectID(userId) {
		json.NewEncoder(res).Encode("User id not found")
	}
	_id, _ := primitive.ObjectIDFromHex(userId)
	user := models.DeleteUser(userController.client, _id)
	if user != false {
		json.NewEncoder(res).Encode("No user found")
		return
	}
	json.NewEncoder(res).Encode(user)
}
