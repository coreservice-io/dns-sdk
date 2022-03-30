package cert

import (
	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

func Apply(applyDomain string, pullZoneName string, hostedDomain string, client *dns_client.Client) (cert string, key string, err error) {
	url := client.EndPoint + "/api/cert/apply/custom"
	postData := commonMsg.CustomDomainCertMsg{
		ApplyDomain:  applyDomain,
		PullZoneName: pullZoneName,
		HostedDomain: hostedDomain,
	}

	var certInfo commonMsg.CertContentResp
	req := httpTools.ApiRequest{
		Result: &certInfo,
	}
	httpTools.POST_(url, client.Token, postData, 120, &req)
	if !req.Ok() {
		return "", "", req.Err
	}
	return certInfo.CertContent, certInfo.KeyContent, nil
}
