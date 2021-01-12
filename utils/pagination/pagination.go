package pagination

import (
	db "github.com/JIeeiroSst/store/component"
	"github.com/JIeeiroSst/store/graph/model"
	bs "github.com/JIeeiroSst/store/utils/base-code"
)

type Edges struct {
	Cursor string
	Node interface{}
}

func Pagination(source []interface{},first int,after string)(*model.PageInfo,interface{},int){
	var Totalcount int64
	db.GetConn().Model(source).Count(&Totalcount)
	id:=bs.DeCode(after)

	db.GetConn().Where("id > ?",id).First(&source)

	return &model.PageInfo{}, nil, 0
}