package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-client/httpTools"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/model"
)

type Client struct {
	EndPoint string
	Token    string
	UserInfo *commonMsg.UserInfoResp
	Domain   *model.Domain
}

func New(token string, domain string, endPoint string) (*Client, error) {
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

	//get domain
	url = endPoint + "/api/domain/querybyname"
	var respDomain model.Domain
	err = httpTools.POST(url, token, &commonMsg.QueryDomainByNameMsg{domain}, 5, &respDomain)
	if err != nil {
		return nil, err
	}
	if respDomain.ID == 0 {
		return nil, errors.New("domain not exist")
	}

	client := &Client{
		EndPoint: endPoint,
		Token:    token,
		UserInfo: &userInfo,
		Domain:   &respDomain,
	}

	return client, nil
}
