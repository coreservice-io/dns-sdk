package record

import (
	"errors"
	"fmt"

	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/model"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

func Add(recordName string, recordType string, ttl uint32, client *dns_client.Client) (*model.Record, error) {
	if recordType != "CNAME" && recordType != "A" {
		return nil, errors.New("only support A and CNAME record")
	}

	url := client.EndPoint + "/api/record/add"
	postData := commonMsg.AddRecordMsg{
		DomainId: client.Domain.ID,
		Name:     recordName,
		Type:     recordType,
		TTL:      ttl,
	}

	var newRecord model.Record
	err := httpTools.POST(url, client.Token, postData, 10, &newRecord)
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func DeleteByRecordName(recordName string, recordType string, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/deletebyname"
	postData := commonMsg.DeleteRecordByNameMsg{
		DomainId:   client.Domain.ID,
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

func ForbiddenByRecordName(recordName string, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/updatebyname"
	postData := commonMsg.UpdateRecordByNameMsg{
		DomainId:   client.Domain.ID,
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

func UpdateByRecordName(recordName string, ttl uint32, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/updatebyname"
	postData := commonMsg.UpdateRecordByNameMsg{
		DomainId:   client.Domain.ID,
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

func QueryByGivenList(recordNameArray []string, client *dns_client.Client) ([]model.Record, error) {
	url := client.EndPoint + "/api/record/querylist"
	postData := commonMsg.QueryRecordListMsg{
		DomainId:       client.Domain.ID,
		RecordNameList: recordNameArray,
	}
	var records []model.Record
	err := httpTools.POST(url, client.Token, postData, 5, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func QueryByNamePattern(namePattern string, limit int, offset int, client *dns_client.Client) (records []*model.Record, totalCount int64, err error) {
	url := client.EndPoint + "/api/record/query"
	postData := commonMsg.QueryRecordMsg{
		DomainId:    client.Domain.ID,
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
