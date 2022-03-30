package record

import (
	"errors"
	"fmt"

	dns_common "github.com/coreservice-io/dns-common"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/data"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

func Add(domain string, recordName string, recordType string, ttl uint32, client *dns_client.Client) (*data.Record, error) {
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

	var newRecord data.Record
	req := httpTools.ApiRequest{
		Result: newRecord,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, req.Err
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
	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, nil)
	if !req.Ok() {
		return req.Err
	}
	return nil
}

func DeleteByRecordId(recordId uint, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/delete/%d", recordId)
	req := httpTools.ApiRequest{}
	httpTools.Get(url, client.Token, nil)
	if !req.Ok() {
		return req.Err
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
	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, nil)
	if !req.Ok() {
		return req.Err
	}
	return nil
}

func ForbiddenByRecordId(recordId uint, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
	postData := commonMsg.UpdateRecordMsg{
		Forbidden: &forbidden,
	}
	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, nil)
	if !req.Ok() {
		return req.Err
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
	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, nil)
	if !req.Ok() {
		return req.Err
	}
	return nil
}

func UpdateByRecordId(recordId uint, ttl uint32, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
	postData := commonMsg.UpdateRecordMsg{
		TTL:       &ttl,
		Forbidden: &forbidden,
	}
	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return req.Err
	}
	return nil
}

func QueryByGivenList(domain string, recordNameArray []string, recordType string, client *dns_client.Client) ([]data.Record, error) {
	url := client.EndPoint + "/api/record/querylist"
	postData := commonMsg.QueryRecordListMsg{
		DomainName:     domain,
		RecordNameList: recordNameArray,
		RecordType:     recordType,
	}
	var records []data.Record
	req := httpTools.ApiRequest{
		Result: &records,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, req.Err
	}
	return records, nil
}

// QueryByNamePattern query records by recordName pattern, recordId, recordType
//  if set namePattern="",recordId=0 or recordType="",query will ignore the condition
func QueryByNamePattern(domain string, namePattern string, recordId uint, recordType string, limit int, offset int, client *dns_client.Client) (records []*data.Record, totalCount int64, e error) {
	url := client.EndPoint + "/api/record/querybydomainname"
	postData := commonMsg.QueryRecordByDomainNameMsg{
		DomainName:  domain,
		NamePattern: namePattern,
		RecordId:    recordId,
		RecordType:  recordType,
		Limit:       limit,
		Offset:      offset,
	}
	var respInfo commonMsg.QueryRecordResp
	req := httpTools.ApiRequest{
		Result: &respInfo,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, 0, req.Err
	}
	return respInfo.Records, respInfo.Count, nil
}
