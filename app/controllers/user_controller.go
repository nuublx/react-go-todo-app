package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nuublx/react-go-todo-app/pkg/utils"
	"github.com/nuublx/react-go-todo-app/platform/hash"
	mongodb "github.com/nuublx/react-go-todo-app/platform/mongo"
	types "github.com/nuublx/react-go-todo-app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(c *fiber.Ctx) error {
	var err error
	var registerModel = new(types.RegisterRequest)
	if err = c.BodyParser(registerModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	err = utils.RegisterRequestValidator(registerModel)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	var filter = bson.M{"$or": []bson.M{
		{"email": registerModel.Email},
		{"username": registerModel.UserName}}}
	var searchResult types.User
	err = mongodb.UsersCollection.FindOne(c.Context(), filter).Decode(&searchResult)
	if err != nil && err != mongo.ErrNoDocuments {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	if err == nil {
		if searchResult.Email == registerModel.Email {
			return c.Status(400).JSON(fiber.Map{"err": true, "msg": "Email already exists"})
		}
		if searchResult.UserName == registerModel.UserName {
			return c.Status(400).JSON(fiber.Map{"err": true, "msg": "Username already exists"})
		}
	}

	hash, salt, err := hash.HashPassword(registerModel.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	var newUser = types.User{
		UserName:    registerModel.UserName,
		Email:       registerModel.Email,
		PhoneNumber: registerModel.PhoneNumber,
		Hash:        hash,
		Salt:        salt,
	}
	result, err := mongodb.UsersCollection.InsertOne(c.Context(), newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}
	var response = new(types.LoginResponse)
	newUser.ID = result.InsertedID.(primitive.ObjectID)
	response.AccessToken, err = utils.GenerateJWT(newUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	return c.Status(201).JSON(response)
}

func Login(c *fiber.Ctx) error {
	var err error
	var loginModel = new(types.LoginRequest)
	if err = c.BodyParser(loginModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	var filter = bson.M{"$or": []bson.M{
		{"email": loginModel.Email},
		{"username": loginModel.UserName}}}
	var user types.User
	err = mongodb.UsersCollection.FindOne(c.Context(), filter).Decode(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": "Invalid credentials"})
	}

	if !hash.CheckPasswordHash(loginModel.Password, user.Salt, user.Hash) {
		return c.Status(400).JSON(fiber.Map{"err": true, "msg": "Invalid credentials"})
	}

	var response = new(types.LoginResponse)
	response.AccessToken, err = utils.GenerateJWT(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"err": true, "msg": err.Error()})
	}

	return c.Status(200).JSON(response)
}
