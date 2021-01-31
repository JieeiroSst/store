// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Node interface {
	IsNode()
}

type FeedBackConnection struct {
	PageInfo   *PageInfo     `json:"pageInfo"`
	Edges      *FeedBackEdge `json:"edges"`
	TotalCount *int          `json:"totalCount"`
}

type FeedBackEdge struct {
	Cursor *string    `json:"cursor"`
	Node   *FeedBacks `json:"node"`
}

type FeedBacks struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Address   *string `json:"address"`
	Content   *string `json:"content"`
	CreatedAt *string `json:"createdAt"`
}

func (FeedBacks) IsNode() {}

type Friend struct {
	ID      string     `json:"id"`
	Friends []*Profile `json:"friends"`
}

func (Friend) IsNode() {}

type InputFeedBacks struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Email   *string `json:"email"`
	Address *string `json:"address"`
	Content *string `json:"content"`
}

type InputMenues struct {
	Text         *string `json:"text"`
	Link         *string `json:"link"`
	DisplayOrder *int    `json:"displayOrder"`
	Target       *string `json:"target"`
}

type InputNewTag struct {
	TagID *int `json:"tagId"`
	NewID *int `json:"newId"`
}

type InputNews struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Detail      *string `json:"detail"`
	CreatedAt   *string `json:"createdAt"`
	TopHot      *string `json:"topHot"`
	ViewCount   *int    `json:"viewCount"`
	Content     *string `json:"content"`
	TagID       *int    `json:"tagId"`
	Active      *bool   `json:"active"`
}

type InputProfile struct {
	UserID     *string `json:"userId"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Address    *string `json:"address"`
	Phone      *string `json:"phone"`
	CreatedAt  *string `json:"createdAt"`
	ModifiedBy *string `json:"modifiedBy"`
}

type InputTags struct {
	Name *string `json:"name"`
}

type Menues struct {
	ID     string  `json:"id"`
	Text   *string `json:"text"`
	Link   *string `json:"link"`
	Target *string `json:"target"`
}

func (Menues) IsNode() {}

type NewTag struct {
	ID    string  `json:"id"`
	TagID *int    `json:"tagId"`
	NewID *int    `json:"newId"`
	Tags  []*Tags `json:"tags"`
	News  []*News `json:"news"`
}

func (NewTag) IsNode() {}

type NewTagsConnection struct {
	PageInfo   *PageInfo      `json:"pageInfo"`
	Edges      []*NewTagsEdge `json:"edges"`
	TotalCount *int           `json:"totalCount"`
}

type NewTagsEdge struct {
	Cursor *string `json:"cursor"`
	Node   *NewTag `json:"node"`
}

type News struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Detail      *string `json:"detail"`
	CreatedAt   *string `json:"createdAt"`
	TopHot      *string `json:"topHot"`
	ViewCount   *int    `json:"viewCount"`
	Content     *string `json:"content"`
	TagID       *int    `json:"tagId"`
	Active      *bool   `json:"active"`
	Tags        []*Tags `json:"tags"`
}

func (News) IsNode() {}

type NewsConnection struct {
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*NewsEdge `json:"edges"`
	TotalCount *int        `json:"totalCount"`
}

type NewsEdge struct {
	Cursor *string `json:"cursor"`
	Node   *News   `json:"node"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

type PaginationInput struct {
	First *int    `json:"first"`
	After *string `json:"after"`
}

type Profile struct {
	ID        string  `json:"id"`
	UserID    *string `json:"userId"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Address   *string `json:"address"`
	Phone     *string `json:"phone"`
	CreatedAt *string `json:"createdAt"`
}

func (Profile) IsNode() {}

type ProfileConnection struct {
	PageInfo   *PageInfo      `json:"pageInfo"`
	Edges      []*ProfileEdge `json:"edges"`
	TotalCount *int           `json:"totalCount"`
}

type ProfileEdge struct {
	Cursor *string  `json:"cursor"`
	Node   *Profile `json:"node"`
}

type ResultCheck struct {
	Status  *bool   `json:"status"`
	Message *string `json:"message"`
}

type ResultFeedBacks struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	Phone     *string `json:"phone"`
	Email     *string `json:"email"`
	Address   *string `json:"address"`
	Content   *string `json:"content"`
	CreatedAt *string `json:"createdAt"`
}

type ResultFriend struct {
	ID string `json:"id"`
}

type ResultMenues struct {
	Text         *string `json:"text"`
	Link         *string `json:"link"`
	DisplayOrder *int    `json:"displayOrder"`
	Target       *string `json:"target"`
}

type ResultNewTag struct {
	ID    string `json:"id"`
	TagID *int   `json:"tagId"`
	NewID *int   `json:"newId"`
}

type ResultNews struct {
	ID          string  `json:"id"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Image       *string `json:"image"`
	Detail      *string `json:"detail"`
	CreatedAt   *string `json:"createdAt"`
	TopHot      *string `json:"topHot"`
	ViewCount   *int    `json:"viewCount"`
	Content     *string `json:"content"`
	TagID       *int    `json:"tagId"`
	Active      *bool   `json:"active"`
}

type ResultProfile struct {
	ID         string  `json:"id"`
	UserID     *string `json:"userId"`
	FirstName  *string `json:"firstName"`
	LastName   *string `json:"lastName"`
	Address    *string `json:"address"`
	Phone      *string `json:"phone"`
	CreatedAt  *string `json:"createdAt"`
	ModifiedBy *string `json:"modifiedBy"`
}

type ResultSystemConfig struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type ResultTags struct {
	Name  *string `json:"name"`
	NewID *int    `json:"newId"`
}

type SystemConfig struct {
	ID    string  `json:"id"`
	Name  *string `json:"name"`
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

func (SystemConfig) IsNode() {}

type SystemConfigConnection struct {
	PageInfo   *PageInfo           `json:"pageInfo"`
	Edges      []*SystemConfigEdge `json:"edges"`
	TotalCount *int                `json:"totalCount"`
}

type SystemConfigEdge struct {
	Cursor *string       `json:"cursor"`
	Node   *SystemConfig `json:"node"`
}

type Tags struct {
	ID   string  `json:"id"`
	Name *string `json:"name"`
}

func (Tags) IsNode() {}

type TagsConnection struct {
	PageInfo   *PageInfo   `json:"pageInfo"`
	Edges      []*TagsEdge `json:"edges"`
	TotalCount *int        `json:"totalCount"`
}

type TagsEdge struct {
	Cursor *string `json:"cursor"`
	Node   *Tags   `json:"node"`
}

type Users struct {
	ID         string   `json:"id"`
	Username   *string  `json:"username"`
	Password   *string  `json:"password"`
	UserDetail *Profile `json:"userDetail"`
}

func (Users) IsNode() {}
