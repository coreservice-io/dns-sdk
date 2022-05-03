package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/common_msg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

type Client struct {
	EndPoint string
	Token    string
	UserInfo *common_msg.User
}

func New(token string, endPoint string) (*Client, error) {
	//get userInfo
	url := endPoint + "/api/user/info"
	var resp common_msg.Msg_Resp_UserInfo
	err := api.Get(url, token, &resp)
	if err != nil {
		return nil, err
	}
	if resp.User.Id == 0 {
		return nil, errors.New("user not exist")
	}

	client := &Client{
		EndPoint: endPoint,
		Token:    token,
		UserInfo: resp.User,
	}

	return client, nil
}
