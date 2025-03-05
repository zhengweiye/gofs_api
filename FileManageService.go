package gofs_api

import "time"

type FileManageService interface {
	GetPageListOfBind(token string, curPage, pageSize int, storeId, appId, envId, bucketId, fileName, fileMd5, businessId string, beginDate, endDate time.Time) (pageData PageData[FileManageBindVo])
	GetPageListOfUnBind(token string, curPage, pageSize int, storeId, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageUnBindVo])
	GetPageListOfTemp(token string, curPage, pageSize int, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageTempVo])
	Delete(token string, ids []string)
	ValidateFile(token string)
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
	fileName, fileMd5, businessId string, beginDate, endDate time.Time) (pageData PageData[FileManageBindVo]) {
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
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	pageData = result.Data
	return
}

func (f FileManageServiceImpl) GetPageListOfUnBind(token string, curPage, pageSize int, storeId, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageUnBindVo]) {
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
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	pageData = result.Data
	return
}

func (f FileManageServiceImpl) GetPageListOfTemp(token string, curPage, pageSize int, fileName, fileMd5 string, beginDate, endDate time.Time) (pageData PageData[FileManageTempVo]) {
	result, err := httpPost[PageData[FileManageTempVo]](f.client, "fileManage/getPageListOfTemp", token, map[string]any{
		"curPage":   curPage,
		"pageSize":  pageSize,
		"fileName":  fileName,
		"fileMd5":   fileMd5,
		"beginDate": beginDate,
		"endDate":   endDate,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	pageData = result.Data
	return
}

func (f FileManageServiceImpl) Delete(token string, ids []string) {
	result, err := httpPost[any](f.client, "fileManage/delete", token, map[string]any{
		"ids": ids,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (f FileManageServiceImpl) ValidateFile(token string) {
	result, err := httpPost[any](f.client, "fileManage/validateFile", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}
