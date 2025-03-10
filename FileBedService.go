package gofs_api

import "errors"

type FileBedService interface {
	GetList(token, parentId string) (list []FileBedVo, err error)
	Mkdir(token, parentDir, dirName string) (err error)
	Rename(token, fileId, fileName string) (err error)
	Move(token, fileId, pid string) (err error)
	Delete(token, id string) (err error)
	Upload(token, pid, fileName string, data []byte) (err error)
}

type FileBedServiceImpl struct {
	client *Client
}

func newFileBedService(c *Client) FileBedService {
	return FileBedServiceImpl{
		client: c,
	}
}

func (f FileBedServiceImpl) GetList(token, parentId string) (list []FileBedVo, err error) {
	result, err := httpPost[[]FileBedVo](f.client, "fileBed/getList", token, map[string]any{
		"parentId": parentId,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	list = result.Data
	return
}

func (f FileBedServiceImpl) Mkdir(token, parentDir, dirName string) (err error) {
	result, err := httpPost[any](f.client, "fileBed/mkdir", token, map[string]any{
		"parentDir": parentDir,
		"dirName":   dirName,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}

func (f FileBedServiceImpl) Rename(token, fileId, fileName string) (err error) {
	result, err := httpPost[any](f.client, "fileBed/rename", token, map[string]any{
		"fileId":   fileId,
		"fileName": fileName,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}

func (f FileBedServiceImpl) Move(token, fileId, pid string) (err error) {
	result, err := httpPost[any](f.client, "fileBed/move", token, map[string]any{
		"fileId": fileId,
		"pid":    pid,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}

func (f FileBedServiceImpl) Delete(token, id string) (err error) {
	result, err := httpPost[any](f.client, "fileBed/delete", token, map[string]any{
		"id": id,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}

func (f FileBedServiceImpl) Upload(token, pid, fileName string, data []byte) (err error) {
	result, err := upload[any](f.client, "fileBed/upload", token, data, fileName, map[string]string{
		"parentDir": pid,
		"fileName":  fileName,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}
