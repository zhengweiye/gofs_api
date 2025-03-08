package gofs_api

type LoginService interface {
	Login(loginAccount, loginPassword string) (login LoginVo)
	Logout(token string)
}

type LoginServiceImpl struct {
	client *Client
}

func newLoginService(c *Client) LoginService {
	return LoginServiceImpl{
		client: c,
	}
}

func (l LoginServiceImpl) Login(loginAccount, loginPassword string) (login LoginVo) {
	result, err := httpPost[LoginVo](l.client, "login", "", map[string]any{
		"loginAccount":  loginAccount,
		"loginPassword": loginPassword,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	login = result.Data
	return
}

func (l LoginServiceImpl) Logout(token string) {
	result, err := httpPost[any](l.client, "logout", token, map[string]any{
		"token": token,
	})
	if err != nil {
		panic(err)
	}
	if result.Code != 200 {
		panic(result.Msg)
	}
	return
}
