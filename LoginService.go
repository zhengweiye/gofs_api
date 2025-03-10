package gofs_api

import "errors"

type LoginService interface {
	Login(loginAccount, loginPassword string) (login LoginVo, err error)
	Logout(token string) (err error)
}

type LoginServiceImpl struct {
	client *Client
}

func newLoginService(c *Client) LoginService {
	return LoginServiceImpl{
		client: c,
	}
}

func (l LoginServiceImpl) Login(loginAccount, loginPassword string) (login LoginVo, err error) {
	result, err := httpPost[LoginVo](l.client, "login", "", map[string]any{
		"loginAccount":  loginAccount,
		"loginPassword": loginPassword,
	})
	if err != nil {
		return
	}
	if result.Code != 200 {
		err = errors.New(result.Msg)
		return
	}
	login = result.Data
	return
}

func (l LoginServiceImpl) Logout(token string) (err error) {
	result, err := httpPost[any](l.client, "logout", token, map[string]any{
		"token": token,
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
