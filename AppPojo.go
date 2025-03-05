package gofs_api

type AppVo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
	SortNum   int    `json:"sortNum"`
}
