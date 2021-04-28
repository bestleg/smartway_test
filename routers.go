package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"smartway_test/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func RemovePerson(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Delete(&person, "people.id = ?", userID)
	if person.ID == 0 {
		http.Error(w, "Person not found or deleted", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}
func GetByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Joins("Company").Find(&person, "person.id = ?", userID)
	//db.Find(&person, "people.id = ?", userID)
	if person.ID == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}

func GetAll(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	persons := []models.Person{}
	db.Joins("Company").Find(&persons)
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(persons)
}

func TestPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", person)
}
