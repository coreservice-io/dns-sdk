package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/common_msg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func (c *Client) ApplyCert(applyDomain string, tag string, hostedDomain string) (cert string, key string, err error) {
	url := c.EndPoint + "/api/cert/apply/custom"
	postData := common_msg.Msg_Req_ApplyCustomCert{
		Apply_domain:  applyDomain,
		Txt_name_tag:  tag,
		Hosted_domain: hostedDomain,
	}

	var resp common_msg.Msg_Resp_ApplyCustomCert
	err = api.POST_(url, c.Token, postData, 120, &resp)
	if err != nil {
		return "", "", err
	}
	if resp.Meta_status < 0 {
		return "", "", errors.New(resp.Meta_message)
	}

	return resp.Cert_content, resp.Key_content, nil
}
