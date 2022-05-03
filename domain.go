package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/common_msg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func (c *Client) QueryDomain(domainName string) (*common_msg.Domain, error) {
	url := c.EndPoint + "/api/domain/query"
	postData := common_msg.Msg_Req_QueryDomain{
		Filter: common_msg.Msg_Req_QueryDomain_Filter{
			Name: &domainName,
		},
		Limit:  1,
		Offset: 0,
	}
	var resp common_msg.Msg_Resp_QueryDomain
	err := api.POST(url, c.Token, postData, &resp)
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
