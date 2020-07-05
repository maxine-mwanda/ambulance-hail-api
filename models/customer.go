package models

import (
	"database/sql"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/utils"
)

func CreateCustomer (data entities.NewCustomer, db *sql.DB) (err error) {
	query := "insert into customer (first_name, last_name, phone_number, gender) values (?,?,?,?)"
	_, err = db.Exec(query, data.FirstName, data.LastName, data.PhoneNumber, data.Gender)
	if err != nil {
		utils.Log("an error occured", err)
		return
	}
	utils.Log("success")
	return
}