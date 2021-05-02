package routers

import (
	"encoding/json"
	"net/http"
	"smartway_test/database"
	"smartway_test/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func PersonRemove(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	idStr := ps.ByName("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		RetError(err, w)
		return
	}

	database.DataBase.Delete(&person, "people.id = ?", userID)
	if person.ID == 0 {
		http.Error(w, "Person deleted or not found", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(person)
}

func PersonAdd(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	person := models.Person{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&person)
	if err != nil {
		RetError(err, w)
		return
	}
	database.DataBase.Create(&person)
	w.Write([]byte(strconv.Itoa(int(person.ID))))
}

func PersonUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	params := models.Person{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&params)
	if err != nil {
		RetError(err, w)
		return
	}
	var person models.Person
	idStr := ps.ByName("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		RetError(err, w)
	}
	database.DataBase.Find(&person, "id = ?", id)
	err = database.DataBase.Model(&person).Updates(params).Error
	if err != nil {
		RetError(err, w)
	}
}

func RetError(err error, w http.ResponseWriter) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
