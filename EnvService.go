package gofs_api

import "errors"

type EnvService interface {
	GetPageList(token string, curPage, pageSize int, appId string) (pageData PageData[EnvVo], err error)
	GetList(token string) (list []Select, err error)
	GetUpdateData(token, id string) (env EnvVo, err error)
	Insert(token, appId, code, name string, sortNum int) (err error)
	Update(token, id, appId, code, name string, sortNum int) (err error)
	Delete(token, id string) (err error)
}

type EnvServiceImpl struct {
	client *Client
}

func newEnvService(c *Client) EnvService {
	return EnvServiceImpl{
		client: c,
	}
}

func (e EnvServiceImpl) GetPageList(token string, curPage, pageSize int, appId string) (pageData PageData[EnvVo], err error) {
	result, err := httpPost[PageData[EnvVo]](e.client, "env/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"appId":    appId,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	pageData = result.Data
	return
}

func (e EnvServiceImpl) GetList(token string) (list []Select, err error) {
	result, err := httpPost[[]Select](e.client, "env/getList", token, map[string]any{})
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

func (e EnvServiceImpl) GetUpdateData(token, id string) (env EnvVo, err error) {
	result, err := httpPost[EnvVo](e.client, "env/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	env = result.Data
	return
}

func (e EnvServiceImpl) Insert(token, appId, code, name string, sortNum int) (err error) {
	result, err := httpPost[any](e.client, "env/insert", token, map[string]any{
		"appId":   appId,
		"code":    code,
		"name":    name,
		"sortNum": sortNum,
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

func (e EnvServiceImpl) Update(token, id, appId, code, name string, sortNum int) (err error) {
	result, err := httpPost[any](e.client, "env/update", token, map[string]any{
		"id":      id,
		"appId":   appId,
		"code":    code,
		"name":    name,
		"sortNum": sortNum,
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

func (e EnvServiceImpl) Delete(token, id string) (err error) {
	result, err := httpPost[any](e.client, "env/delete", token, map[string]any{
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
