package models

import (
	"database/sql"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/utils"
)

func CreateDriver (data entities.NewDriver, db *sql.DB) (err error) {
	query := "insert into driver (first_name, last_name, phone_number, id_no, gender, password, salt) values (?,?,?,?,?,?,?)"
	salt := "salt"
	_, err = db.Exec(query, data.FirstName, data.LastName, data.PhoneNumber,data.IdNumber, data.Gender, data.Password, salt)
	if err != nil {
		utils.Log("an error occured", err)
		return
	}
	utils.Log("success")
	return
}
