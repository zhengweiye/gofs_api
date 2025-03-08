package gofs_api

type Client struct {
	http        string
	ip          string
	port        int
	contextPath string
}

type Option struct {
	Http        string
	Ip          string
	Port        int
	ContextPath string
}

func Create(opt Option) *Client {
	return &Client{
		http:        opt.Http,
		ip:          opt.Ip,
		port:        opt.Port,
		contextPath: opt.ContextPath,
	}
}

func (c *Client) GetLoginService() LoginService {
	return newLoginService(c)
}

func (c *Client) GetAppService() AppService {
	return newAppService(c)
}

func (c *Client) GetBucketService() BucketService {
	return newBucketService(c)
}

func (c *Client) GetDustbinService() DustbinService {
	return newDustbinService(c)
}

func (c *Client) GetEnvService() EnvService {
	return newEnvService(c)
}

func (c *Client) GetFileBedService() FileBedService {
	return newFileBedService(c)
}

func (c *Client) GetFileManageService() FileManageService {
	return newFileManageService(c)
}

func (c *Client) GetStorageService() StorageService {
	return newStorageService(c)
}
