package domain

import (
	"errors"

	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func Query(domainName string, client *dns_client.Client) (*commonMsg.Domain, error) {
	url := client.EndPoint + "/api/domain/query"
	postData := commonMsg.Msg_Req_QueryDomain{
		Filter: commonMsg.Msg_Req_QueryDomain_Filter{
			Name: &domainName,
		},
		Limit:  1,
		Offset: 0,
	}
	var resp commonMsg.Msg_Resp_QueryDomain
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	if len(resp.Domain_list) == 0 {
		return nil, errors.New("domain not exist")
	}

	return resp.Domain_list[0], nil
}
