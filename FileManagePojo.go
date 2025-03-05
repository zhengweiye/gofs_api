package gofs_api

type FileManageBindVo struct {
	Id          string `json:"id"`
	FileName    string `json:"fileName"`
	FileMd5     string `json:"fileMd5"`
	AppName     string `json:"appName"`
	EnvName     string `json:"envName"`
	BucketName  string `json:"bucketName"`
	ServerName  string `json:"serverName"`
	StorePath   string `json:"storePath"`
	FilePath    string `json:"filePath"`
	FileType    string `json:"fileType"`
	CreateTime  string `json:"createTime"`
	AllowDelete bool   `json:"allowDelete"`
	StoreId     string `json:"-"`
	AppId       string `json:"-"`
	EnvId       string `json:"-"`
	BucketId    string `json:"-"`
}

type FileManageUnBindVo struct {
	Id          string `json:"id"`
	FileName    string `json:"fileName"`
	FileMd5     string `json:"fileMd5"`
	ServerName  string `json:"serverName"`
	BucketName  string `json:"bucketName"`
	StorePath   string `json:"storePath"`
	FilePath    string `json:"filePath"`
	FileType    string `json:"fileType"`
	CreateTime  string `json:"createTime"`
	AllowDelete bool   `json:"allowDelete"`
	StoreId     string `json:"-"`
	BucketId    string `json:"-"`
}

type FileManageTempVo struct {
	Id          string `json:"id"`
	FileName    string `json:"fileName"`
	FileMd5     string `json:"fileMd5"`
	StorePath   string `json:"storePath"`
	FilePath    string `json:"filePath"`
	FileType    string `json:"fileType"`
	CreateTime  string `json:"createTime"`
	AllowDelete bool   `json:"allowDelete"`
	StoreId     string `json:"-"`
}
