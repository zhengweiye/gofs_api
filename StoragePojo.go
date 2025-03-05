package gofs_api

type StorageListVo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	IsDefault    bool   `json:"isDefault"`
	NginxDomain  string `json:"nginxDomain"`
	NginxPrefix  string `json:"nginxPrefix"`
	Ip           string `json:"ip"`
	Port         int    `json:"port"`
	LoginAccount string `json:"loginAccount"`
	Volumes      string `json:"volumes"`
	TotalSize    string `json:"totalSize"`
	UsedSize     string `json:"usedSize"`
	SortNum      int    `json:"sortNum"`
}

type StorageVo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	IsDefault    bool   `json:"isDefault"`
	NginxDomain  string `json:"nginxDomain"`
	NginxPrefix  string `json:"nginxPrefix"`
	Ip           string `json:"ip"`
	Port         int    `json:"port"`
	LoginAccount string `json:"loginAccount"`
	LoginPwd     string `json:"loginPwd"`
	Volumes      string `json:"volumes"`
	SortNum      int    `json:"sortNum"`
}
