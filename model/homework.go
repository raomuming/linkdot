package model

import (
	"github.com/globalsign/mgo/bson"
)

type Homework struct {
	Id		bson.ObjectId 	`bson:"_id" json:"id"`
	Name 	string 			`bson:"name" json:"name"`
}

const (
	db 	= "Linkdot"
	collection = "Homework"
)

func (h *Homework) Insert(homework Homework) error {
	return Insert(db, collection, homework)
}

func (h *Homework) FindAll() ([]Homework, error) {
	var result []Homework
	err := FindAll(db, collection, nil, nil, &result)
	return result, err
}

func (h *Homework) FindById(id string) (Homework, error) {
	var result Homework
	err := FindOne(db, collection, bson.M{"_id": bson.ObjectIdHex(id)}, nil, &result)
	return result, err
}

func (h *Homework) Update(homework Homework) error {
	return Update(db, collection, bson.M{"_id": homework.Id}, homework)
}

func (h *Homework) Remove(id string) error {
	return Remove(db, collection, bson.M{"_id": bson.ObjectIdHex(id)})
}