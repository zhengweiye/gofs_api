package gofs_api

type StorageService interface {
	GetPageList(token string, curPage, pageSize int, name string) (pageData PageData[StorageListVo])
	GetList(token string) (list []Select)
	GetUpdateData(token, id string) (storage StorageVo)
	Insert(token, name string, isDefault bool, nginxDomain, nginxPrefix, ip string, port int, loginAccount, loginPwd, volumes string, sortNum int)
	Update(token, id, name string, isDefault bool, nginxDomain, nginxPrefix, ip string, port int, loginAccount, volumes string, sortNum int)
	Delete(token, id string)
	UpdatePwd(token, id, loginPwd string)
}

type StorageServiceImpl struct {
	client *Client
}

func newStorageService(c *Client) StorageService {
	return StorageServiceImpl{
		client: c,
	}
}

func (s StorageServiceImpl) GetPageList(token string, curPage, pageSize int, name string) (pageData PageData[StorageListVo]) {
	result, err := httpPost[PageData[StorageListVo]](s.client, "storage/getPageList", token, map[string]any{
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

func (s StorageServiceImpl) GetList(token string) (list []Select) {
	result, err := httpPost[[]Select](s.client, "storage/getList", token, map[string]any{})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	list = result.Data
	return
}

func (s StorageServiceImpl) GetUpdateData(token, id string) (storage StorageVo) {
	result, err := httpPost[StorageVo](s.client, "storage/getUpdateData", token, map[string]any{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	storage = result.Data
	return
}

func (s StorageServiceImpl) Insert(token, name string, isDefault bool, nginxDomain, nginxPrefix, ip string, port int, loginAccount, loginPwd, volumes string, sortNum int) {
	result, err := httpPost[any](s.client, "storage/insert", token, map[string]any{
		"name":         name,
		"isDefault":    isDefault,
		"nginxDomain":  nginxDomain,
		"nginxPrefix":  nginxPrefix,
		"ip":           ip,
		"port":         port,
		"loginAccount": loginAccount,
		"loginPwd":     loginPwd,
		"volumes":      volumes,
		"sortNum":      sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (s StorageServiceImpl) Update(token, id, name string, isDefault bool, nginxDomain, nginxPrefix, ip string, port int, loginAccount, volumes string, sortNum int) {
	result, err := httpPost[any](s.client, "storage/update", token, map[string]any{
		"id":           id,
		"name":         name,
		"isDefault":    isDefault,
		"nginxDomain":  nginxDomain,
		"nginxPrefix":  nginxPrefix,
		"ip":           ip,
		"port":         port,
		"loginAccount": loginAccount,
		"volumes":      volumes,
		"sortNum":      sortNum,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}

func (s StorageServiceImpl) Delete(token, id string) {
	result, err := httpPost[any](s.client, "storage/delete", token, map[string]any{
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

func (s StorageServiceImpl) UpdatePwd(token, id, loginPwd string) {
	result, err := httpPost[any](s.client, "storage/updatePwd", token, map[string]any{
		"id":       id,
		"loginPwd": loginPwd,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}
