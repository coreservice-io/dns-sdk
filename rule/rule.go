package rule

import (
	"errors"
	"fmt"

	"github.com/coreservice-io/dns-common/commonMsg"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/tools/api"
)

func AddRuleByRecordName(domain string, recordName string, recordType string, version int, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) (*commonMsg.Rule, error) {
	url := client.EndPoint + "/api/rule/addbyrecordname"
	postData := commonMsg.Msg_Req_AddRuleByRecordName{
		DomainName:    domain,
		RecordName:    recordName,
		RecordType:    recordType,
		SysVersion:    version,
		ContinentCode: continentCode,
		CountryCode:   countryCode,
		StartTime:     startTime,
		EndTime:       endTime,
		Destination:   dest,
		Weight:        weight,
	}

	var resp commonMsg.Msg_Resp_RuleInfo
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return &resp.Rule, nil
}

//func AddRuleByRecordId(recordId uint, version int, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) (*commonMsg.Rule, error) {
//	url := client.EndPoint + "/api/rule/add"
//	postData := commonMsg.Msg_Req_AddRule{
//		RecordId:      recordId,
//		SysVersion:    version,
//		ContinentCode: continentCode,
//		CountryCode:   countryCode,
//		StartTime:     startTime,
//		EndTime:       endTime,
//		Destination:   dest,
//		Weight:        weight,
//	}
//
//	var resp commonMsg.Msg_Resp_RuleInfo
//	err := api.POST(url, client.Token, postData, &resp)
//	if err != nil {
//		return nil, err
//	}
//	if resp.Meta_status < 0 {
//		return nil, errors.New(resp.Meta_message)
//	}
//
//	return &resp.Rule, nil
//}

func QueryRulesByRecordName(domain string, recordName string, recordType string, client *dns_client.Client) ([]*commonMsg.Rule, error) {
	url := client.EndPoint + "/api/rule/querybyrecordname"
	postData := commonMsg.Msg_Req_QueryRulesByRecordName{
		DomainName: domain,
		RecordName: recordName,
		RecordType: recordType,
	}
	var resp commonMsg.Msg_Resp_Rules
	err := api.POST(url, client.Token, postData, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Meta_status < 0 {
		return nil, errors.New(resp.Meta_message)
	}

	return resp.Rules, nil
}

//func QueryRulesByRecordId(recordId uint, client *dns_client.Client) ([]*commonMsg.Rule, error) {
//	url := client.EndPoint + fmt.Sprintf("/api/rule/query/%d", recordId)
//	var resp commonMsg.Msg_Resp_Rules
//	err := api.Get(url, client.Token, &resp)
//	if err != nil {
//		return nil, err
//	}
//	if resp.Meta_status < 0 {
//		return nil, errors.New(resp.Meta_message)
//	}
//
//	return resp.Rules, nil
//}

func Delete(ruleId uint, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/rule/delete/%d", ruleId)
	var resp api.API_META_STATUS
	err := api.Get(url, client.Token, &resp)
	if err != nil {
		return err
	}
	if resp.Meta_status < 0 {
		return errors.New(resp.Meta_message)
	}

	return nil
}

func UpdateRule(ruleId uint, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/rule/update/%d", ruleId)
	postData := commonMsg.Msg_Req_UpdateRule{
		ContinentCode: &continentCode,
		CountryCode:   &countryCode,
		StartTime:     &startTime,
		EndTime:       &endTime,
		Destination:   &dest,
		Weight:        &weight,
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
