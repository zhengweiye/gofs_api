package gofs_api

type DustbinListVo struct {
	Id         string `json:"id"`
	FileName   string `json:"fileName"`
	FileMd5    string `json:"fileMd5"`
	FileType   string `json:"fileType"`
	ServerName string `json:"serverName"`
	StorePath  string `json:"storePath"`
	FilePath   string `json:"filePath"`
	CreateTime string `json:"createTime"`
	DeleteTime string `json:"deleteTime"`
}
