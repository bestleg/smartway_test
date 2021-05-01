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

func ShowPersonsFromCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := []models.Person{}
	idStr := ps.ByName("id")
	companyID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Joins("Department").Joins("Passport").Find(&person, "people.company_id = ?", companyID)
	if len(person) == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}

func ShowPersonsFromDepartment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	department := models.Department{}
	person := []models.Person{}
	depName := ps.ByName("name")

	db.Find(&department, "departments.name = ?", depName)
	db.Joins("Department").Joins("Passport").Find(&person, "department_id = ?", department.ID)

	if department.ID == 0 {
		http.Error(w, "Department not found", http.StatusNotFound)
		return
	}

	if len(person) == 0 {
		http.Error(w, "Department is empty", http.StatusNotFound)
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
	db.Find(&person, "people.id = ?", userID)
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
