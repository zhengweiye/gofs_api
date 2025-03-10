package gofs_api

import (
	"errors"
	"time"
)

type FileManageService interface {
	GetPageListOfBind(token string, curPage, pageSize int, storeId, appId, envId, bucketId, fileName, fileMd5, businessId string, beginDate, endDate time.Time) (pageData PageData[FileManageBindVo], err error)
	GetPageListOfUnBind(token string, curPage, pageSize int, storeId, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageUnBindVo], err error)
	GetPageListOfTemp(token string, curPage, pageSize int, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageTempVo], err error)
	Delete(token string, ids []string) (err error)
	ValidateFile(token string) (err error)
}

type FileManageServiceImpl struct {
	client *Client
}

func newFileManageService(c *Client) FileManageService {
	return FileManageServiceImpl{
		client: c,
	}
}

func (f FileManageServiceImpl) GetPageListOfBind(token string, curPage, pageSize int, storeId, appId, envId, bucketId,
	fileName, fileMd5, businessId string, beginDate, endDate time.Time) (pageData PageData[FileManageBindVo], err error) {
	result, err := httpPost[PageData[FileManageBindVo]](f.client, "fileManage/getPageListOfBind", token, map[string]any{
		"curPage":    curPage,
		"pageSize":   pageSize,
		"storeId":    storeId,
		"appId":      appId,
		"envId":      envId,
		"bucketId":   bucketId,
		"fileName":   fileName,
		"fileMd5":    fileMd5,
		"businessId": businessId,
		"beginDate":  beginDate,
		"endDate":    endDate,
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

func (f FileManageServiceImpl) GetPageListOfUnBind(token string, curPage, pageSize int, storeId, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageUnBindVo], err error) {
	result, err := httpPost[PageData[FileManageUnBindVo]](f.client, "fileManage/getPageListOfUnBind", token, map[string]any{
		"curPage":   curPage,
		"pageSize":  pageSize,
		"storeId":   storeId,
		"fileName":  fileName,
		"fileMd5":   fileMd5,
		"beginDate": beginDate,
		"endDate":   endDate,
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

func (f FileManageServiceImpl) GetPageListOfTemp(token string, curPage, pageSize int, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageTempVo], err error) {
	result, err := httpPost[PageData[FileManageTempVo]](f.client, "fileManage/getPageListOfTemp", token, map[string]any{
		"curPage":   curPage,
		"pageSize":  pageSize,
		"fileName":  fileName,
		"fileMd5":   fileMd5,
		"beginDate": beginDate,
		"endDate":   endDate,
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

func (f FileManageServiceImpl) Delete(token string, ids []string) (err error) {
	result, err := httpPost[any](f.client, "fileManage/delete", token, map[string]any{
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

func (f FileManageServiceImpl) ValidateFile(token string) (err error) {
	result, err := httpPost[any](f.client, "fileManage/validateFile", token, map[string]any{})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	return
}
