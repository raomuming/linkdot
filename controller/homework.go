package controller

import (
	"encoding/json"
	"net/http"

	"github.com/raomuming/linkdot/model"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

var (
	dao = model.Homework{}
)

