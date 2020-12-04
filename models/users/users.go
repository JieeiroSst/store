package users

import "time"

type Users struct {
	Id int
	Username string
	Password string
	FirstName string
	LastName string
	Address string
	Phone string
	CreatedAt time.Time
	CreatedBy string
	ModifiedDate time.Time
	ModifiedBy string
	Status int
}

