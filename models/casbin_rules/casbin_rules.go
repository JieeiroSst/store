package casbin_rules

import "github.com/JIeeiroSst/store/models/paginations"

type CasbinRules struct {
	Id int
	PType string
	V0 string
	V1 string
	V2 string
}

type CasbinRuleEdges struct {
	Cursor string
	Node CasbinRules
}

type CasbinConnection struct {
	PageInfo paginations.PageInfo
	Edges [] CasbinRuleEdges
	TotalCount int
}

type InputCasbinRule struct {
	Id int
	PType string
	V0 string
	V1 string
	V2 string
}