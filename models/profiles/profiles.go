package profiles

import (
	"time"
	"github.com/JIeeiroSst/store/models/paginations"
)

type Profiles struct {
	Id int
	friendId int
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

type ProfileEdge struct {
	Cursor string
	Node Profiles
}

type ProfileConnection struct {
	PageInfo paginations.PageInfo
	Edges []ProfileEdge
	TotalCount int
}

type InputProfile struct {
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