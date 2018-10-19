package model

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id     bson.ObjectId `bosn:"_id" json:"id"`
	Name   string        `bson:"name" json:"name"`
	OpenId string        `bson:"openId" json:"openId"`
}

type WechatModel struct {
	Code        string `json:"code"`
	CryptedData string `json:"crypted_data"`
	Iv          string `json:"iv"`
}

const (
	userCollection = "User"
)

func (u *User) Insert(user User) error {
	return Insert(db, userCollection, user)
}

func (u *User) FindAll() ([]User, error) {
	var result []User
	err := FindAll(db, userCollection, nil, nil, &result)
	return result, err
}

func (u *User) FindById(id string) (User, error) {
	var result User
	err := FindOne(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (u *User) FindByOpenId(openId string) (User, error) {
	var result User
	err := FindOne(db, userCollection, bson.M{"openId": openId}, nil, &result)
	return result, err
}

func (u *User) Update(user User) error {
	return Update(db, userCollection, bson.M{"_id": user.Id}, user)
}

func (u *User) Remove(id string) error {
	return Remove(db, userCollection, bson.M{"_id": bson.ObjectIdHex(id)})
}
