CREATE TABLE customer (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
 first_name VARCHAR (30) NOT NULL,
 last_name VARCHAR (30) NOT NULL,
 phone_number VARCHAR (20) NOT NULL UNIQUE,
 gender VARCHAR (20) NOT NULL,
 created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
 updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE driver (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
first_name VARCHAR (30) NOT NULL,
last_name VARCHAR (30) NOT NULL,
phone_number VARCHAR (20) NOT NULL UNIQUE,
id_no VARCHAR (20) NOT NULL UNIQUE,
gender VARCHAR (20) NOT NULL,
password VARCHAR (30) NOT NULL,
salt VARCHAR (30),
created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE ambulance (
ambulance_id INT (15) PRIMARY KEY AUTO_INCREMENT,
car_model VARCHAR (30) NOT NULL,
number_plate VARCHAR (30) NOT NULL UNIQUE,
active bool
);

CREATE TABLE customer_requests (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
customer INT,
ambulance INT,
time_requested DATETIME DEFAULT CURRENT_TIMESTAMP,
status VARCHAR (30),
end_time DATETIME,
FOREIGN KEY (customer) references customer(user_id),
FOREIGN KEY (ambulance) references ambulance(ambulance_id)
);

CREATE TABLE driver_duty (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
driver INT,
ambulance INT,
start_time DATETIME DEFAULT CURRENT_TIMESTAMP,
end_time DATETIME,
FOREIGN KEY (driver) references driver(user_id),
FOREIGN KEY (ambulance) references ambulance(ambulance_id)
);

CREATE TABLE admin (
user_id INT (15) PRIMARY KEY AUTO_INCREMENT,
user_name VARCHAR (30),
password VARCHAR (30),
salt VARCHAR (30)
);