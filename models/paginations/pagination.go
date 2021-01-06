package paginations

type PaginationInput struct {
	First int `json:"first"`
	After string `json:"after"` 
}

type PageInfo struct {
	HashNextPage bool
	HasPreviousPage bool
	StartCursor string
	EndCursor string
}