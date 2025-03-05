package gofs_api

type FileBedVo struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FilePath string `json:"filePath"`
	IsDir    bool   `json:"isDir"`
}
