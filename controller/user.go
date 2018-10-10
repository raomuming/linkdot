package controller

import (
	"encoding/json"
	"net/http"

	"github.com/raomuming/linkdot/auth"
	"github.com/raomuming/linkdot/model"
	"github.com/globalsign/mgo/bson"
)

const (
	db = "Linkdot"
	collection = "User"
)

func Register(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {
	
}