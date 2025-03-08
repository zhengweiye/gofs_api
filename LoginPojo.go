package gofs_api

type LoginVo struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Token    string `json:"token"`
}
