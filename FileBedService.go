package gofs_api

type FileBedService interface {
	GetList(token, parentId string) (list []FileBedVo)
	Mkdir(token, parentDir, dirName string)
	Rename(token, fileId, fileName string)
	Move(token, fileId, pid string)
	Delete(token, id string)
	Upload(token, pid, fileName string, data []byte)
}

type FileBedServiceImpl struct {
	client *Client
}

func newFileBedService(c *Client) FileBedService {
	return FileBedServiceImpl{
		client: c,
	}
}

func (f FileBedServiceImpl) GetList(token, parentId string) (list []FileBedVo) {
	result, err := httpPost[[]FileBedVo](f.client, "fileBed/getList", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	list = result.Data
	return
}

func (f FileBedServiceImpl) Mkdir(token, parentDir, dirName string) {
	result, err := httpPost[any](f.client, "fileBed/mkdir", token, map[string]any{
		"parentDir": parentDir,
		"dirName":   dirName,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileBedServiceImpl) Rename(token, fileId, fileName string) {
	result, err := httpPost[any](f.client, "fileBed/rename", token, map[string]any{
		"fileId":   fileId,
		"fileName": fileName,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileBedServiceImpl) Move(token, fileId, pid string) {
	result, err := httpPost[any](f.client, "fileBed/move", token, map[string]any{
		"fileId": fileId,
		"pid":    pid,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileBedServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](f.client, "fileBed/delete", token, map[string]any{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileBedServiceImpl) Upload(token, pid, fileName string, data []byte) {
	result, err := upload[any](f.client, "fileBed/upload", token, data, fileName, map[string]string{
		"pid":      pid,
		"fileName": fileName,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}
