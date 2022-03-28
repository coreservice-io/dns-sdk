package rule

import (
	"errors"
	"strconv"

	dns_client "github.com/coreservice-io/dns-client"
	"github.com/coreservice-io/dns-client/httpTools"
	"github.com/coreservice-io/dns-client/record"
	"github.com/coreservice-io/dns-common/commonMsg"
	"github.com/coreservice-io/dns-common/model"
)

func Add(recordName string, version int, continentCode string, countryCode string, startTime string, endTime string, dest string, weight int, client *dns_client.Client) (*model.Rule, error) {
	recordInfo, err := record.QueryByGivenList([]string{recordName}, client)
	if err != nil {
		return nil, err
	}
	if len(recordInfo) == 0 {
		return nil, errors.New("record not exist")
	}

	url := client.EndPoint + "/api/rule/add"
	postData := commonMsg.AddRuleMsg{
		SysVersion:    version,
		RecordId:      recordInfo[0].ID,
		ContinentCode: continentCode,
		CountryCode:   countryCode,
		StartTime:     startTime,
		EndTime:       endTime,
		Destination:   dest,
		Weight:        weight,
	}

	var newRule model.Rule
	err = httpTools.POST(url, client.Token, postData, 10, &newRule)
	if err != nil {
		return nil, err
	}

	return &newRule, nil
}

func QueryRules(recordName string, client *dns_client.Client) ([]model.Rule, error) {
	recordInfo, err := record.QueryByGivenList([]string{recordName}, client)
	if err != nil {
		return nil, err
	}
	if len(recordInfo) == 0 {
		return nil, errors.New("record not exist")
	}

	recordIdStr := strconv.FormatUint(uint64(recordInfo[0].ID), 10)

	url := client.EndPoint + "/api/rule/query/" + recordIdStr
	var newRule []model.Rule
	err = httpTools.Get(url, client.Token, 10, &newRule)
	if err != nil {
		return nil, err
	}
	return newRule, nil
}

func Delete() {

}

func UpdateSysRule() {

}
