package gofs_api

import "errors"

type DustbinService interface {
	GetPageList(token string, curPage, pageSize int, storeId, fileName string) (pageData PageData[DustbinListVo], err error)
	Delete(token string, ids []string) (err error)
}

type DustbinServiceImpl struct {
	client *Client
}

func newDustbinService(c *Client) DustbinService {
	return DustbinServiceImpl{
		client: c,
	}
}

func (d DustbinServiceImpl) GetPageList(token string, curPage, pageSize int, storeId, fileName string) (pageData PageData[DustbinListVo], err error) {
	result, err := httpPost[PageData[DustbinListVo]](d.client, "fileDustbin/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"storeId":  storeId,
		"fileName": fileName,
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

func (d DustbinServiceImpl) Delete(token string, ids []string) (err error) {
	result, err := httpPost[any](d.client, "fileDustbin/delete", token, map[string]any{
		"ids": ids,
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
