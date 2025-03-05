package gofs_api

type EnvService interface {
	GetPageList(token string, curPage, pageSize int, appId string) (pageData PageData[EnvVo])
	GetList(token string) (list []Select)
	GetUpdateData(token, id string) (env EnvVo)
	Insert(token, appId, code, name string, sortNum int)
	Update(token, id, appId, code, name string, sortNum int)
	Delete(token, id string)
}

type EnvServiceImpl struct {
	client *Client
}

func newEnvService(c *Client) EnvService {
	return EnvServiceImpl{
		client: c,
	}
}

func (e EnvServiceImpl) GetPageList(token string, curPage, pageSize int, appId string) (pageData PageData[EnvVo]) {
	result, err := httpPost[PageData[EnvVo]](e.client, "env/getPageList", token, map[string]any{
		"curPage":  curPage,
		"pageSize": pageSize,
		"appId":    appId,
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

func (e EnvServiceImpl) GetList(token string) (list []Select) {
	result, err := httpPost[[]Select](e.client, "env/getList", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	list = result.Data
	return
}

func (e EnvServiceImpl) GetUpdateData(token, id string) (env EnvVo) {
	result, err := httpPost[EnvVo](e.client, "env/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	env = result.Data
	return
}

func (e EnvServiceImpl) Insert(token, appId, code, name string, sortNum int) {
	result, err := httpPost[any](e.client, "env/insert", token, map[string]any{
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

func (e EnvServiceImpl) Update(token, id, appId, code, name string, sortNum int) {
	result, err := httpPost[any](e.client, "env/update", token, map[string]any{
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

func (e EnvServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](e.client, "env/delete", token, map[string]any{
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
