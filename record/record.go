package record

import (
	"errors"
	"fmt"

	dns_common "github.com/coreservice-io/dns-common"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/model"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

func Add(domain string, recordName string, recordType string, ttl uint32, client *dns_client.Client) (*model.Record, error) {
	if recordType != dns_common.TypeCNAME && recordType != dns_common.TypeA {
		return nil, errors.New("only support A and CNAME record")
	}

	url := client.EndPoint + "/api/record/addbydomainname"
	postData := commonMsg.AddRecordByDomainNameMsg{
		DomainName: domain,
		Name:       recordName,
		Type:       recordType,
		TTL:        ttl,
	}

	var newRecord model.Record
	err := httpTools.POST(url, client.Token, postData, 10, &newRecord)
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func DeleteByRecordName(domain string, recordName string, recordType string, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/deletebyname"
	postData := commonMsg.DeleteRecordByNameMsg{
		DomainName: domain,
		RecordName: recordName,
		RecordType: recordType,
	}
	err := httpTools.POST(url, client.Token, postData, 10, nil)
	if err != nil {
		return err
	}
	return nil
}

func DeleteByRecordId(recordId uint, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/delete/%d", recordId)
	err := httpTools.Get(url, client.Token, 10, nil)
	if err != nil {
		return err
	}
	return nil
}

func ForbiddenByRecordName(domain string, recordName string, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/updatebyname"
	postData := commonMsg.UpdateRecordByNameMsg{
		DomainName: domain,
		RecordName: recordName,
		Forbidden:  &forbidden,
	}
	err := httpTools.POST(url, client.Token, postData, 5, nil)
	if err != nil {
		return err
	}
	return nil
}

func ForbiddenByRecordId(recordId uint, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
	postData := commonMsg.UpdateRecordMsg{
		Forbidden: &forbidden,
	}
	err := httpTools.POST(url, client.Token, postData, 5, nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateByRecordName(domain string, recordName string, ttl uint32, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/updatebyname"
	postData := commonMsg.UpdateRecordByNameMsg{
		DomainName: domain,
		RecordName: recordName,
		TTL:        &ttl,
		Forbidden:  &forbidden,
	}
	err := httpTools.POST(url, client.Token, postData, 5, nil)
	if err != nil {
		return err
	}
	return nil
}

func UpdateByRecordId(recordId uint, ttl uint32, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
	postData := commonMsg.UpdateRecordMsg{
		TTL:       &ttl,
		Forbidden: &forbidden,
	}
	err := httpTools.POST(url, client.Token, postData, 5, nil)
	if err != nil {
		return err
	}
	return nil
}

func QueryByGivenList(domain string, recordNameArray []string, client *dns_client.Client) ([]model.Record, error) {
	url := client.EndPoint + "/api/record/querylist"
	postData := commonMsg.QueryRecordListMsg{
		DomainName:     domain,
		RecordNameList: recordNameArray,
	}
	var records []model.Record
	err := httpTools.POST(url, client.Token, postData, 5, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func QueryByNamePattern(domain string, namePattern string, limit int, offset int, client *dns_client.Client) (records []*model.Record, totalCount int64, err error) {
	url := client.EndPoint + "/api/record/querybydomainname"
	postData := commonMsg.QueryRecordByDomainNameMsg{
		DomainName:  domain,
		NamePattern: namePattern,
		Limit:       limit,
		Offset:      offset,
	}
	var respInfo commonMsg.QueryRecordResp
	err = httpTools.POST(url, client.Token, postData, 5, &respInfo)
	if err != nil {
		return nil, 0, err
	}
	return respInfo.Records, respInfo.Count, nil
}
