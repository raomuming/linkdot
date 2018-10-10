package model

type User struct {
	UserName string `bson:"username" json:"username"`
}