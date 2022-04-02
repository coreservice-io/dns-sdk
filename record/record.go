package record

import (
	"errors"

	dns_common "github.com/coreservice-io/dns-common"
	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func Add(domain string, recordName string, recordType string, ttl uint32, client *dns_client.Client) (*commonMsg.Record, error) {
	if recordType != dns_common.TypeCNAME && recordType != dns_common.TypeA {
		return nil, errors.New("only support A and CNAME record")
	}

	url := client.EndPoint + "/api/record/add_by_domain_name"
	postData := commonMsg.Msg_Req_AddRecordByDomainName{
		Domain_name: domain,
		Name:        recordName,
		Type:        recordType,
		TTL:         ttl,
	}

	var resp commonMsg.Msg_Resp_RecordInfo
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return &resp.Record, nil
}

func DeleteByRecordName(domain string, recordName string, recordType string, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/delete_by_record_name"
	postData := commonMsg.Msg_Req_DeleteRecordByName{
		Domain_name: domain,
		Record_name: recordName,
		Record_type: recordType,
	}
	var resp api.API_META_STATUS
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

//func DeleteByRecordId(recordId uint, client *dns_client.Client) error {
//	url := client.EndPoint + fmt.Sprintf("/api/record/delete/%d", recordId)
//	var resp api.API_META_STATUS
//	err := api.Get(url, client.Token, &resp)
//	if err != nil {
//		return err
//	}
//	if resp.Meta_status < 0 {
//		return errors.New(resp.Meta_message)
//	}
//	return nil
//}

func ForbiddenByRecordName(domain string, recordName string, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/update_by_record_name"
	postData := commonMsg.Msg_Req_UpdateRecordByName{
		Domain_name: domain,
		Record_name: recordName,
		Forbidden:   &forbidden,
	}
	var resp api.API_META_STATUS
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

//func ForbiddenByRecordId(recordId uint, forbidden bool, client *dns_client.Client) error {
//	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
//	postData := commonMsg.Msg_Req_UpdateRecord{
//		Forbidden: &forbidden,
//	}
//	var resp api.API_META_STATUS
//	err := api.POST(url, client.Token, postData, &resp)
//	if err != nil {
//		return err
//	}
//	if resp.Meta_status < 0 {
//		return errors.New(resp.Meta_message)
//	}
//	return nil
//}

func UpdateByRecordName(domain string, recordName string, ttl uint32, forbidden bool, client *dns_client.Client) error {
	url := client.EndPoint + "/api/record/update_by_record_name"
	postData := commonMsg.Msg_Req_UpdateRecordByName{
		Domain_name: domain,
		Record_name: recordName,
		TTL:         &ttl,
		Forbidden:   &forbidden,
	}
	var resp api.API_META_STATUS
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

//func UpdateByRecordId(recordId uint, ttl uint32, forbidden bool, client *dns_client.Client) error {
//	url := client.EndPoint + fmt.Sprintf("/api/record/update/%d", recordId)
//	postData := commonMsg.Msg_Req_UpdateRecord{
//		TTL:       &ttl,
//		Forbidden: &forbidden,
//	}
//	var resp api.API_META_STATUS
//	err := api.POST(url, client.Token, postData, &resp)
//	if err != nil {
//		return err
//	}
//	if resp.Meta_status < 0 {
//		return errors.New(resp.Meta_message)
//	}
//	return nil
//}

func QueryByGivenList(domain string, recordNameArray []string, recordType string, client *dns_client.Client) ([]*commonMsg.Record, error) {
	url := client.EndPoint + "/api/record/query_by_given_name"
	postData := commonMsg.Msg_Req_QueryRecordByGivenName{
		Domain_name:      domain,
		Record_name_list: recordNameArray,
		Record_type:      recordType,
	}
	var resp commonMsg.Msg_Resp_QueryRecordByGivenName
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Records, nil
}

// QueryByNamePattern query records by recordName pattern, recordId, recordType
//  if set namePattern="",recordId=0 or recordType="",query will ignore the condition
func QueryByNamePattern(domain string, namePattern string, recordId uint, recordType string, limit int, offset int, client *dns_client.Client) (records []*commonMsg.Record, totalCount int64, e error) {
	url := client.EndPoint + "/api/record/query_by_domain_name"
	postData := commonMsg.Msg_Req_QueryRecordByDomainName{
		Domain_name:  domain,
		Name_pattern: namePattern,
		Record_id:    recordId,
		Record_type:  recordType,
		Limit:        limit,
		Offset:       offset,
	}
	var resp commonMsg.Msg_Resp_QueryRecord
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, 0, err
	}
	if resp.Meta_status < 0 {
		return nil, 0, errors.New(resp.Meta_message)
	}

	return resp.Records, resp.Count, nil
}
