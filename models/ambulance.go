package models

import (
	"database/sql"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/utils"
)

func CreateAmbulance (data entities.NewAmbulance, db *sql.DB) (err error) {
	query := "insert into ambulance (car_model, number_plate, active) values (?,?,?)"
	_, err = db.Exec(query, data.CarModel, data.NumberPlate, data.Active)
	if err != nil {
		utils.Log("an error occured", err)
		return
	}
	utils.Log("success")
	return
}

