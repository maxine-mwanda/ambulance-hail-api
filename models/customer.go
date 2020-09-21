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

func ListCustomers (db *sql.DB) (Customers []entities.Customer, err error){
	query:= "select user_id, first_name, last_name, phone_number, id_no, gender,created_at from driver;"
	rows, err := db.Query(query)
	if err!=nil {
		utils.Log("an error occured", err)
		return
	}
	for rows.Next(){
		var customer entities.Customer
		err=rows.Scan(&customer.UserID, &customer.FirstName, &customer.LastName, &customer.PhoneNumber, &customer.Gender, &customer.CreatedAt)
		if err !=nil {
			utils.Log("and error occured", err)
			continue
		}

		//customer = append(Customers, customer)
	}
	return
}