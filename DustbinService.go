package gofs_api

type DustbinService interface {
	GetPageList(token string, curPage, pageSize int, storeId, fileName string) (pageData PageData[DustbinListVo])
	Delete(token, id string)
}

type DustbinServiceImpl struct {
	client *Client
}

func newDustbinService(c *Client) DustbinService {
	return DustbinServiceImpl{
		client: c,
	}
}

func (d DustbinServiceImpl) GetPageList(token string, curPage, pageSize int, storeId, fileName string) (pageData PageData[DustbinListVo]) {
	result, err := httpPost[PageData[DustbinListVo]](d.client, "fileDustbin/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"storeId":  storeId,
		"fileName": fileName,
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

func (d DustbinServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](d.client, "fileDustbin/delete", token, map[string]any{
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
