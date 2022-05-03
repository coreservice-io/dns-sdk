package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/common_msg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

type NewRuleData struct {
	ContinentCode string
	CountryCode   string
	Destination   string
	Weight        int64
}

func (c *Client) AddRule(domainName string, recordName string, recordType string, rules []*NewRuleData) ([]*common_msg.Rule, error) {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, errors.New("record not exsit")
	}

	newRules := []*common_msg.Rule{}
	url := c.EndPoint + "/api/rule/add"
	for _, v := range rules {
		postData := common_msg.Msg_Req_AddRule{
			Record_id:      records[0].Id,
			Continent_code: v.ContinentCode,
			Country_code:   v.CountryCode,
			Destination:    v.Destination,
			Weight:         v.Weight,
		}

		var resp common_msg.Msg_Resp_AddRule
		err := api.POST(url, c.Token, postData, &resp)
		if err != nil {
			return nil, err
		}
		if resp.Meta_status < 0 {
			return nil, errors.New(resp.Meta_message)
		}
		newRules = append(newRules, resp.Rule)
	}

	return newRules, nil
}

func (c *Client) QueryRules(domainName string, recordName string, recordType string) ([]*common_msg.Rule, error) {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, errors.New("record not exsit")
	}

	url := c.EndPoint + "/api/rule/query"
	postData := common_msg.Msg_Req_QueryRule{
		Filter: common_msg.Msg_Req_QueryRule_Filter{
			Record_id: records[0].Id,
		},
	}
	var resp common_msg.Msg_Resp_QueryRules
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Rules, nil
}

func (c *Client) DeleteRules(ids []int64) error {
	url := c.EndPoint + "/api/rule/delete"
	for _, v := range ids {
		postData := common_msg.Msg_Req_DeleteRule{
			Filter: common_msg.Msg_Req_DeleteRule_Filter{
				Id: []int64{v},
			},
		}
		var resp common_msg.API_META_STATUS
		err := api.POST(url, c.Token, postData, &resp)
		if err != nil {
			return err
		}
		if resp.Meta_status < 0 {
			return errors.New(resp.Meta_message)
		}
	}
	return nil
}

func (c *Client) DeleteAllRules(domainName string, recordName string, recordType string) error {
	rules, err := c.QueryRules(domainName, recordName, recordType)
	if err != nil {
		return err
	}
	if len(rules) == 0 {
		return nil
	}

	ids := []int64{}
	for _, v := range rules {
		ids = append(ids, v.Id)
	}
	return c.DeleteRules(ids)
}
