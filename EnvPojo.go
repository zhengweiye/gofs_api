package gofs_api

type EnvVo struct {
	Id      string `json:"id"`
	AppId   string `json:"appId"`
	AppName string `json:"appName"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	SortNum int    `json:"sortNum"`
}
