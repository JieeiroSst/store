package models

import "time"

type Menues struct {
	ID           int  `json:"id"`
	Text         string `json:"text"`
	Link         string `json:"link"`
	DisplayOrder int    `json:"displayOrder"`
	Target       string `json:"target"`
}

type FeedBacks struct {
	ID        *int  `json:"id"`
	Name      *string `json:"name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Address   *string `json:"address"`
	Content   *string `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
}


type NewTag struct {
	ID    *int  `json:"id"`
	TagId *int    `json:"tagId"`
	NewId *int
}

type News struct {
	ID          *int  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Detail      *string `json:"detail"`
	CreatedAt   *time.Time `json:"createdAt"`
	ViewCount   *int    `json:"viewCount"`
	Content     *string `json:"content"`
	TagID       *int    `json:"tagId"`
	Active      *bool   `json:"active"`
	Tags        []*Tags `json:"tags"`
}

type Profile struct {
	ID         *int  `json:"id"`
	UserID     *int `json:"userId"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Address    *string `json:"address"`
	Phone      *string `json:"phone"`
	CreatedAt  *time.Time `json:"createdAt"`
}

type SystemConfig struct {
	ID    *int  `json:"id"`
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type Users struct {
	ID         int   `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Permission string `json:"permission"`
	UserDetail Profile `json:"userDetail"`
}

type Tags struct {
	ID     *int    `json:"id"`
	Name   *string   `json:"name"`
}
