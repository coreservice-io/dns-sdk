package rule

import (
	"fmt"

	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/data"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/httpTools"
)

func AddRuleByRecordName(domain string, recordName string, recordType string, version int, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) (*data.Rule, error) {
	url := client.EndPoint + "/api/rule/addbyrecordname"
	postData := commonMsg.AddRuleByRecordNameMsg{
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

	var newRule data.Rule
	req := httpTools.ApiRequest{
		Result: &newRule,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, req.Err
	}

	return &newRule, nil
}

func AddRuleByRecordId(recordId uint, version int, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) (*data.Rule, error) {
	url := client.EndPoint + "/api/rule/add"
	postData := commonMsg.AddRuleMsg{
		RecordId:      recordId,
		SysVersion:    version,
		ContinentCode: continentCode,
		CountryCode:   countryCode,
		StartTime:     startTime,
		EndTime:       endTime,
		Destination:   dest,
		Weight:        weight,
	}

	var newRule data.Rule
	req := httpTools.ApiRequest{
		Result: &newRule,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, req.Err
	}

	return &newRule, nil
}

func QueryRulesByRecordName(domain string, recordName string, recordType string, client *dns_client.Client) ([]data.Rule, error) {
	url := client.EndPoint + "/api/rule/querybyrecordname"
	postData := commonMsg.QueryRulesByRecordNameMsg{
		DomainName: domain,
		RecordName: recordName,
		RecordType: recordType,
	}
	var rules []data.Rule
	req := httpTools.ApiRequest{
		Result: &rules,
	}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return nil, req.Err
	}
	return rules, nil
}

func QueryRulesByRecordId(recordId uint, client *dns_client.Client) ([]data.Rule, error) {
	url := client.EndPoint + fmt.Sprintf("/api/rule/query/%d", recordId)
	var rules []data.Rule
	req := httpTools.ApiRequest{
		Result: &rules,
	}
	httpTools.Get(url, client.Token, &req)
	if !req.Ok() {
		return nil, req.Err
	}
	return rules, nil
}

func Delete(ruleId uint, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/rule/delete/%d", ruleId)
	req := httpTools.ApiRequest{}
	httpTools.Get(url, client.Token, nil)
	if !req.Ok() {
		return req.Err
	}
	return nil
}

func UpdateRule(ruleId uint, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) error {
	url := client.EndPoint + fmt.Sprintf("/api/rule/update/%d", ruleId)
	postData := commonMsg.UpdateRuleMsg{
		ContinentCode: &continentCode,
		CountryCode:   &countryCode,
		StartTime:     &startTime,
		EndTime:       &endTime,
		Destination:   &dest,
		Weight:        &weight,
	}

	req := httpTools.ApiRequest{}
	httpTools.POST(url, client.Token, postData, &req)
	if !req.Ok() {
		return req.Err
	}

	return nil
}
