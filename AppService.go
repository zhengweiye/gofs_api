package gofs_api

type AppService interface {
	/**
	 * 分页
	 */
	GetPageList(token string, curPage, pageSize int, name string) (pageData PageData[AppVo])

	/**
	 * 下拉框
	 */
	GetList(token string) (list []Select)

	/**
	 * 修改--回显
	 */
	GetUpdateData(token, id string) (app AppVo)

	/**
	 * 新增
	 */
	Insert(token, name, appKey, appSecret string, sortNum int)

	/**
	 * 修改--保存
	 */
	Update(token, id, name, appKey, appSecret string, sortNum int)

	/**
	 * 删除
	 */
	Delete(token, id string)
}

type AppServiceImpl struct {
	client *Client
}

func newAppService(c *Client) AppService {
	return AppServiceImpl{
		client: c,
	}
}

func (a AppServiceImpl) GetPageList(token string, curPage, pageSize int, name string) (pageData PageData[AppVo]) {
	result, err := httpPost[PageData[AppVo]](a.client, "app/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"name":     name,
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

func (a AppServiceImpl) GetList(token string) (list []Select) {
	result, err := httpPost[[]Select](a.client, "app/getList", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	list = result.Data
	return
}

func (a AppServiceImpl) GetUpdateData(token, id string) (app AppVo) {
	result, err := httpPost[AppVo](a.client, "app/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	app = result.Data
	return
}

func (a AppServiceImpl) Insert(token, name, appKey, appSecret string, sortNum int) {
	result, err := httpPost[any](a.client, "app/insert", token, map[string]any{
		"name":      name,
		"appKey":    appKey,
		"appSecret": appSecret,
		"sortNum":   sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (a AppServiceImpl) Update(token, id, name, appKey, appSecret string, sortNum int) {
	result, err := httpPost[any](a.client, "app/update", token, map[string]any{
		"id":        id,
		"name":      name,
		"appKey":    appKey,
		"appSecret": appSecret,
		"sortNum":   sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (a AppServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](a.client, "app/delete", token, map[string]any{
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
