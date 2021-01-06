package system_config

import "github.com/JIeeiroSst/store/models/paginations"

type SystemConfig struct {
	Id int
	Name string
	Types string
	Value string
	Status int
}

type SystemConfigEdge struct {
	Cursor string
	Node SystemConfig
}

type SystemConfigConnection struct {
	PageInfo paginations.PageInfo
	Edges []SystemConfigEdge
	TotalCount int
}