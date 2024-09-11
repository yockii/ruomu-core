package server

import "github.com/yockii/ruomu-core/database"

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type TimeCondition struct {
	Start database.DateTime `json:"start,omitempty" query:"start"`
	End   database.DateTime `json:"end,omitempty" query:"end"`
}

type Paginate struct {
	Total  int64         `json:"total,omitempty"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
	Items  []interface{} `json:"items,omitempty"`
}
