package gofs_api

import "errors"

type BucketService interface {
	GetPageList(token string, curPage, pageSize int, appId, keyWord string) (pageData PageData[BucketListVo], err error)
	GetList(token string) (list []Select, err error)
	GetUpdateData(token, id string) (bucket BucketVo, err error)
	Insert(token, appId, code, name string, sortNum int) (err error)
	Update(token, id, appId, code, name string, sortNum int) (err error)
	Delete(token, id string) (err error)
}

type BucketServiceImpl struct {
	client *Client
}

func newBucketService(c *Client) BucketService {
	return BucketServiceImpl{
		client: c,
	}
}

func (b BucketServiceImpl) GetPageList(token string, curPage, pageSize int, appId, keyWord string) (pageData PageData[BucketListVo], err error) {
	result, err := httpPost[PageData[BucketListVo]](b.client, "bucket/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"appId":    appId,
		"keyWord":  keyWord,
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

func (b BucketServiceImpl) GetList(token string) (list []Select, err error) {
	result, err := httpPost[[]Select](b.client, "bucket/getList", token, map[string]any{})
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

func (b BucketServiceImpl) GetUpdateData(token, id string) (bucket BucketVo, err error) {
	result, err := httpPost[BucketVo](b.client, "bucket/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	bucket = result.Data
	return
}

func (b BucketServiceImpl) Insert(token, appId, code, name string, sortNum int) (err error) {
	result, err := httpPost[any](b.client, "bucket/insert", token, map[string]any{
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

func (b BucketServiceImpl) Update(token, id, appId, code, name string, sortNum int) (err error) {
	result, err := httpPost[any](b.client, "bucket/update", token, map[string]any{
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

func (b BucketServiceImpl) Delete(token, id string) (err error) {
	result, err := httpPost[any](b.client, "bucket/delete", token, map[string]any{
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
