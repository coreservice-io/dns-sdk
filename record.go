package dns_client

import (
	"errors"

	dns_common "github.com/coreservice-io/dns-common"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func (c *Client) AddRecord(domainName string, recordName string, recordType string, ttl uint32) (*commonMsg.Record, error) {
	if recordType != dns_common.TypeCNAME && recordType != dns_common.TypeA {
		return nil, errors.New("only support A and CNAME record")
	}

	//get domain id
	domainInfo, err := c.QueryDomain(domainName)
	if err != nil {
		return nil, err
	}

	url := c.EndPoint + "/api/record/add"
	postData := commonMsg.Msg_Req_AddRecord{
		Domain_id: domainInfo.Id,
		Name:      recordName,
		Type:      recordType,
		TTL:       ttl,
	}

	var resp commonMsg.Msg_Resp_AddRecord
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Record, nil
}

func (c *Client) DeleteRecord(domainName string, recordName string, recordType string) error {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("record not exsit")
	}

	url := c.EndPoint + "/api/record/delete"
	postData := commonMsg.Msg_Req_DeleteRecord{
		Filter: commonMsg.Msg_Req_DeleteRecord_Filter{
			Id: []int64{records[0].Id},
		},
	}
	var resp api.API_META_STATUS
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

func (c *Client) UpdateRecord(domainName string, recordName string, recordType string, forbidden *bool, ttl *uint32) error {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return errors.New("record not exsit")
	}

	url := c.EndPoint + "/api/record/update"
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
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}
	return nil
}

func (c *Client) QueryRecord(domainName string, recordNameArray []string, recordType string, limit int, offset int) (records []*commonMsg.Record, totalCount int64, err error) {
	records = []*commonMsg.Record{}
	//get domain id
	domainInfo, err := c.QueryDomain(domainName)
	if err != nil {
		return
	}

	url := c.EndPoint + "/api/record/query"
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
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return
	}
	if resp.Meta_status < 0 {
		err = errors.New(resp.Meta_message)
		return
	}

	return resp.Records, resp.Count, nil
}
