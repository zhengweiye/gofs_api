package gofs_api

type BucketListVo struct {
	Id      string `json:"id"`
	AppName string `json:"appName"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	SortNum int    `json:"sortNum"`
}

type BucketVo struct {
	Id      string `json:"id"`
	AppId   string `json:"appId"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	SortNum int    `json:"sortNum"`
}
