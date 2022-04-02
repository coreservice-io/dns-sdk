package cert

import (
	"errors"

	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func Apply(applyDomain string, pullZoneName string, hostedDomain string, client *dns_client.Client) (cert string, key string, err error) {
	url := client.EndPoint + "/api/cert/apply/custom"
	postData := commonMsg.Msg_Req_ApplyCustomCert{
		Apply_domain:   applyDomain,
		Pull_zone_name: pullZoneName,
		Hosted_domain:  hostedDomain,
	}

	var resp commonMsg.Msg_Resp_CertContent
	err = api.POST_(url, client.Token, postData, 120, &resp)
	if err != nil {
		return "", "", err
	}
	if resp.Meta_status < 0 {
		return "", "", errors.New(resp.Meta_message)
	}

	return resp.Cert_content, resp.Key_content, nil
}
