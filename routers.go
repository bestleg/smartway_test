package main

import (
	"encoding/json"
	"net/http"
	"smartway_test/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func personRemove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Delete(&person, "people.id = ?", userID)
	if person.ID == 0 {
		http.Error(w, "Person deleted or not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}

func showPersonsFromCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := []models.Person{}
	idStr := ps.ByName("id")
	companyID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Joins("Department").Joins("Passport").Find(&person, "people.company_id = ?", companyID)
	if len(person) == 0 {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}

func showPersonsFromDepartment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func personAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&person)
	if err != nil {
		retError(err, w)
		return
	}
	db.Create(&person)
	w.Write([]byte(strconv.Itoa(int(person.ID))))
}

func personUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := models.Person{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&params)
	if err != nil {
		retError(err, w)
		return
	}
	var curper models.Person
	var id int
	idStr := ps.ByName("id")
	id, _ = strconv.Atoi(idStr)
	db.Find(&curper, "id = ?", id)
	db.Model(&curper).Updates(params)
	db.Find(&curper, "id = ?", id)
	enc := json.NewEncoder(w)
	enc.Encode(&curper)
}

func retError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
