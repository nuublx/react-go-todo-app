package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserName    string             `json:"username"`
	Email       string             `json:"email"`
	PhoneNumber string             `json:"phonenumber"`
	Hash        string             `json:"hash"`
	Salt        string             `json:"salt"`
}

type RegisterRequest struct {
	UserName        string `json:"username"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phonenumber"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmpassword"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"accesstoken"`
}

type ForgetPasswordRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}
