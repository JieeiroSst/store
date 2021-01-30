package models

type Menues struct {
	ID           int  `json:"id"`
	Text         string `json:"text"`
	Link         string `json:"link"`
	DisplayOrder int    `json:"displayOrder"`
	Target       string `json:"target"`
}

type NewTag struct {
	ID    int  `json:"id"`
	TagID int    `json:"tagId"`
	Name  string `json:"name"`
}

type News struct {
	ID              int  `json:"id"`
	Title           string `json:"title"`
	MetaTitle       string `json:"metaTitle"`
	Description     string `json:"description"`
	Image           string `json:"image"`
	CategoryID      int    `json:"categoryId"`
	Detail          string `json:"detail"`
	CreatedAt       string `json:"createdAt"`
	CreatedBy       string `json:"createdBy"`
	ModifiedData    string `json:"modifiedData"`
	ModifiedBy      string `json:"modifiedBy"`
	MetaKeyWord     string `json:"metaKeyWord"`
	MetaDescription string `json:"metaDescription"`
	TopHot          string `json:"topHot"`
	ViewCount       int    `json:"viewCount"`
	Content         string `json:"content"`
	TagID           int    `json:"tagId"`
	Active          bool   `json:"active"`
	Tags            []*Tags `json:"tags"`
}

type Profile struct {
	ID           int  `json:"id"`
	UserID       string `json:"userId"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    string `json:"createdBy"`
	ModifiedDate string `json:"modifiedDate"`
	ModifiedBy   string `json:"modifiedBy"`
	FriendID     []*int  `json:"friendId"`
}

type SystemConfig struct {
	ID    int  `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Users struct {
	ID         int   `json:"id"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	UserDetail Profile `json:"userDetail"`
}

type Tags struct {
	ID     int    `json:"id"`
	Name   string   `json:"name"`
	NewID  int      `json:"newId"`
	News   []News   `json:"news"`
	NewTag []NewTag `json:"newTag"`
}
