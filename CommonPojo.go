package gofs_api

type Result[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

type PageData[T any] struct {
	List          []T   `json:"list"`          // 响应内容
	TotalElements int64 `json:"totalElements"` // 总记录数
	TotalPage     int64 `json:"totalPage"`     // 总页数
}

type Select struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
