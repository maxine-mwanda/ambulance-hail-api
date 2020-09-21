package entities

type NewAdmin struct{
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

//todo: check validity

type NewAmbulance struct {
CarModel string `json:"car_model"`
NumberPlate string `json:"number_plate"`
Active string `json:"active"`
}

type NewCustomer struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
}

type NewDriver struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	IdNumber string `json:"id_no"`
	Gender string `json:"gender"`
	Password string `json:"password"`
}

type Driver struct {
	UserID int `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	IdNumber string `json:"id_no"`
	Gender string `json:"gender"`
	CreatedAt string `json:"created_at"`
}


type Customer struct {
	UserID int `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
	CreatedAt string `json:"created_at"`
}