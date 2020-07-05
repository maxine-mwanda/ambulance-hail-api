package models

import (
	"database/sql"
	"github.com/innv8/ambulance-hail-api/entities"
	"github.com/innv8/ambulance-hail-api/utils"
)

func CreateAdmin (data entities.NewAdmin, db *sql.DB) (err error) {
	query := "insert into admin (user_name, password, salt) values (?,?,?)"
	salt := "salt"
	_, err = db.Exec(query, data.UserName, data.Password,
		salt)
	if err != nil {
		utils.Log("an error occured", err)
		return
	}
	utils.Log("success")
	return
}
