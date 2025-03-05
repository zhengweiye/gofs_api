package gofs_api

type BucketService interface {
	GetPageList(token string, curPage, pageSize int, appId, keyWord string) (pageData PageData[BucketListVo])
	GetList(token string) (list []Select)
	GetUpdateData(token, id string) (bucket BucketVo)
	Insert(token, appId, code, name string, sortNum int)
	Update(token, id, appId, code, name string, sortNum int)
	Delete(token, id string)
}

type BucketServiceImpl struct {
	client *Client
}

func newBucketService(c *Client) BucketService {
	return BucketServiceImpl{
		client: c,
	}
}

func (b BucketServiceImpl) GetPageList(token string, curPage, pageSize int, appId, keyWord string) (pageData PageData[BucketListVo]) {
	result, err := httpPost[PageData[BucketListVo]](b.client, "bucket/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"appId":    appId,
		"keyWord":  keyWord,
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

func (b BucketServiceImpl) GetList(token string) (list []Select) {
	result, err := httpPost[[]Select](b.client, "bucket/getList", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	list = result.Data
	return
}

func (b BucketServiceImpl) GetUpdateData(token, id string) (bucket BucketVo) {
	result, err := httpPost[BucketVo](b.client, "bucket/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	bucket = result.Data
	return
}

func (b BucketServiceImpl) Insert(token, appId, code, name string, sortNum int) {
	result, err := httpPost[any](b.client, "bucket/insert", token, map[string]any{
		"appId":   appId,
		"code":    code,
		"name":    name,
		"sortNum": sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (b BucketServiceImpl) Update(token, id, appId, code, name string, sortNum int) {
	result, err := httpPost[any](b.client, "bucket/update", token, map[string]any{
		"id":      id,
		"appId":   appId,
		"code":    code,
		"name":    name,
		"sortNum": sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (b BucketServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](b.client, "bucket/delete", token, map[string]any{
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
