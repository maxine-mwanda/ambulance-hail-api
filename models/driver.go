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

func ListDrivers (db *sql.DB) (drivers []entities.Driver, err error){
	query:= "select user_id, first_name, last_name, phone_number, id_no, gender,created_at from driver;"
	rows, err := db.Query(query)
	if err!=nil {
		utils.Log("an error occured", err)
		return
	}
	for rows.Next(){
		var driver entities.Driver
		err=rows.Scan(&driver.UserID, &driver.FirstName, &driver.LastName, &driver.PhoneNumber, &driver.IdNumber, &driver.Gender, &driver.CreatedAt)
		if err !=nil {
			utils.Log("and error occured", err)
			continue
		}

		drivers = append(drivers, driver)
	}
	return
}
