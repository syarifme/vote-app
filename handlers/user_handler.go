package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/syarifme/vote-app/models"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {

	user := models.User{
		Model: gorm.Model{
			ID: 1,
		},
		Name:     "Muhammad Syarif",
		Username: "syarifme",
		Email:    "muhsya44@gmail.com",
	}

	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func AddUser(w http.ResponseWriter, r *http.Request) {

}
