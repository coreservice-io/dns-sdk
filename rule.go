package dns_client

import (
	"errors"

	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

type NewRuleData struct {
	ContinentCode string
	CountryCode   string
	Destination   string
	Weight        int64
}

func (c *Client) AddRule(domainName string, recordName string, recordType string, rules []*NewRuleData) ([]*commonMsg.Rule, error) {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, errors.New("record not exsit")
	}

	newRules := []*commonMsg.Rule{}
	url := c.EndPoint + "/api/rule/add"
	for _, v := range rules {
		postData := commonMsg.Msg_Req_AddRule{
			Record_id:      records[0].Id,
			Continent_code: v.ContinentCode,
			Country_code:   v.CountryCode,
			Destination:    v.Destination,
			Weight:         v.Weight,
		}

		var resp commonMsg.Msg_Resp_AddRule
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

func (c *Client) QueryRules(domainName string, recordName string, recordType string) ([]*commonMsg.Rule, error) {
	//get record id
	records, _, err := c.QueryRecord(domainName, []string{recordName}, recordType, 1, 0)
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		return nil, errors.New("record not exsit")
	}

	url := c.EndPoint + "/api/rule/query"
	postData := commonMsg.Msg_Req_QueryRule{
		Filter: commonMsg.Msg_Req_QueryRule_Filter{
			Record_id: records[0].Id,
		},
	}
	var resp commonMsg.Msg_Resp_QueryRules
	err = api.POST(url, c.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Rules, nil
}

func (c *Client) DeleteRules(domainName string, recordName string, recordType string) error {
	rules, err := c.QueryRules(domainName, recordName, recordType)
	if err != nil {
		return err
	}
	if len(rules) == 0 {
		return nil
	}

	url := c.EndPoint + "/api/rule/delete"
	for _, v := range rules {
		postData := commonMsg.Msg_Req_DeleteRule{
			Filter: commonMsg.Msg_Req_DeleteRule_Filter{
				Id: []int64{v.Id},
			},
		}
		var resp commonMsg.API_META_STATUS
		err = api.POST(url, c.Token, postData, &resp)
		if err != nil {
			return err
		}
		if resp.Meta_status < 0 {
			return errors.New(resp.Meta_message)
		}
	}
	return nil
}
