package record

import (
	"errors"

	dns_common "github.com/coreservice-io/dns-common"
	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	domainMgr "github.com/coreservice-io/dns-sdk/domain"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func Add(domainName string, recordName string, recordType string, ttl uint32, client *dns_client.Client) (*commonMsg.Record, error) {
	if recordType != dns_common.TypeCNAME && recordType != dns_common.TypeA {
		return nil, errors.New("only support A and CNAME record")
	}

	//get domain id
	domainInfo, err := domainMgr.Query(domainName, client)
	if err != nil {
		return nil, err
	}

	url := client.EndPoint + "/api/record/add"
	postData := commonMsg.Msg_Req_AddRecord{
		Domain_id: domainInfo.Id,
		Name:      recordName,
		Type:      recordType,
		TTL:       ttl,
	}

	var resp commonMsg.Msg_Resp_AddRecord
	err = api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Record, nil
}

func Delete(domainName string, recordName string, recordType string, client *dns_client.Client) error {
	//get record id
	records, _, err := Query(domainName, []string{recordName}, recordType, 1, 0, client)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("record not exsit")
	}

	url := client.EndPoint + "/api/record/delete"
	postData := commonMsg.Msg_Req_DeleteRecord{
		Filter: commonMsg.Msg_Req_DeleteRecord_Filter{
			Id: []int64{records[0].Id},
		},
	}
	var resp api.API_META_STATUS
	err = api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

func Update(domainName string, recordName string, recordType string, forbidden *bool, ttl *uint32, client *dns_client.Client) error {
	//get record id
	records, _, err := Query(domainName, []string{recordName}, recordType, 1, 0, client)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("record not exsit")
	}

	url := client.EndPoint + "/api/record/update"
	postData := commonMsg.Msg_Req_UpdateRecord{
		Filter: commonMsg.Msg_Req_UpdateRecord_Filter{
			Id: []int64{},
		},
		Update: commonMsg.Msg_Req_UpdateRecord_To{
			TTL:       ttl,
			Forbidden: forbidden,
		},
	}
	var resp api.API_META_STATUS
	err = api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

func Query(domainName string, recordNameArray []string, recordType string, limit int, offset int, client *dns_client.Client) (records []*commonMsg.Record, totalCount int64, err error) {
	records = []*commonMsg.Record{}
	//get domain id
	domainInfo, err := domainMgr.Query(domainName, client)
	if err != nil {
		return
	}

	url := client.EndPoint + "/api/record/query"
	postData := commonMsg.Msg_Req_QueryRecord{
		Filter: commonMsg.Msg_Req_QueryRecord_Filter{
			Domain_id: &domainInfo.Id,
			Name:      &recordNameArray,
			Type:      &recordType,
		},
		Limit:  limit,
		Offset: offset,
	}
	var resp commonMsg.Msg_Resp_QueryRecord
	err = api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return
	}
	if resp.Meta_status < 0 {
		err = errors.New(resp.Meta_message)
		return
	}

	return resp.Records, resp.Count, nil
}
