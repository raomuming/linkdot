package controller

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/raomuming/linkdot/model"
)

var (
	homeworkDao = model.Homework{}
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllHomeworks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var homeworks []model.Homework
	homeworks, err := homeworkDao.FindAll()
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, homeworks)
}

func FindHomework(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := homeworkDao.FindById(id)
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, result)
}

func CreateHomework(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var homework model.Homework

	if err := json.NewDecoder(r.Body).Decode(&homework); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if userId := r.Context().Value("Id"); userId != nil {
		homework.Creator = userId.(string)
	} else {
		responseWithJson(w, http.StatusBadRequest, "not authorized")
		return
	}

	homework.Id = bson.NewObjectId()
	if err := homeworkDao.Insert(homework); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusCreated, homework)
}

func UpdateHomework(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var params model.Homework
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := homeworkDao.Update(params); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteHomework(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := homeworkDao.Remove(id); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
