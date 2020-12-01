package users

import "time"

type Users struct {
	id int
	username string
	password string
	firstName string
	lastName string
	address string
	phone string
	createdAt time.Time
	createdBy string
	modifiedDate time.Time
	modifiedBy string
	status int
}
