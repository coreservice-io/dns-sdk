package cert

import (
	dns_client "github.com/coreservice-io/dns-client"
	"github.com/coreservice-io/dns-client/httpTools"
	"github.com/coreservice-io/dns-common/commonMsg"
)

func Apply(applyDomain string, pullZoneName string, client *dns_client.Client) (cert string, key string, err error) {
	url := client.EndPoint + "/api/cert/apply/custom"
	postData := commonMsg.CustomDomainCertMsg{
		ApplyDomain:  applyDomain,
		PullZoneName: pullZoneName,
		HostedDomain: client.Domain.Name,
	}

	var certInfo commonMsg.CertContentResp
	err = httpTools.POST(url, client.Token, postData, 120, &certInfo)
	if err != nil {
		return "", "", err
	}
	return certInfo.CertContent, certInfo.KeyContent, nil
}
