package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

type Client struct {
	EndPoint string
	Token    string
	UserInfo *commonMsg.UserInfoResp
}

func New(token string, endPoint string) (*Client, error) {
	//get userInfo
	url := endPoint + "/api/user/info"
	var userInfo commonMsg.UserInfoResp
	err := httpTools.Get(url, token, 5, &userInfo)
	if err != nil {
		return nil, err
	}
	if userInfo.ID == 0 {
		return nil, errors.New("user not exist")
	}

	client := &Client{
		EndPoint: endPoint,
		Token:    token,
		UserInfo: &userInfo,
	}

	return client, nil
}
