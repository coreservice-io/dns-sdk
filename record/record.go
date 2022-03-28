package record

import (
	"errors"

	dns_client "github.com/coreservice-io/dns-client"
	"github.com/coreservice-io/dns-client/httpTools"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/model"
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
	err := httpTools.POST(url, client.Token, postData, 5, &newRecord)
	if err != nil {
		return nil, err
	}

	return &newRecord, nil
}

func Delete(recordName string, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/deletebyname"
	postData := commonMsg.DeleteRecordByNameMsg{
		DomainId:   client.Domain.ID,
		RecordName: recordName,
	}
	err := httpTools.POST(url, client.Token, postData, 5, nil)
	if err != nil {
		return err
	}
	return nil
}

func Forbidden(recordName string, forbidden bool, client *dns_client.Client) error {
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

func Update(recordName string, ttl uint32, forbidden bool, client *dns_client.Client) error {
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
