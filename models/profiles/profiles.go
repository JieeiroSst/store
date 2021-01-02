package profiles

import "time"

type Profiles struct {
	UserId int `json:"user_id"`
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
