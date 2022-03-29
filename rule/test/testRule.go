package main

import (
	"log"

	dns_common "github.com/coreservice-io/dns-common"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/rule"
	"github.com/coreservice-io/ipGeo"
)

var token = "cxwinggdadlzdhpcyktmikjj"
var domain = "mesoncdn.com"
var endPoint = "http://127.0.0.1:9001"

func AddRuleByRecordName() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	newRule, err := rule.AddRuleByRecordName(domain, "pullzone1", dns_common.TypeCNAME, 0, ipGeo.AllContinent, ipGeo.AllCountry, dns_common.DayStart, dns_common.DayEnd, "www.google.com", 100, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("newRule", newRule)
	}
}

func AddRuleByRecordId() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	newRule, err := rule.AddRuleByRecordId(9, 0, ipGeo.AllContinent, ipGeo.AllCountry, dns_common.DayStart, dns_common.DayEnd, "www.youtube.com", 100, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("newRule", newRule)
	}
}

func QueryRuleByRecordName() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	rules, err := rule.QueryRulesByRecordName(domain, "pullzone1", dns_common.TypeCNAME, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("rules")
		for _, v := range rules {
			log.Println(v)
		}
	}
}

func QueryRuleByRecordId() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	rules, err := rule.QueryRulesByRecordId(9, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("rules")
		for _, v := range rules {
			log.Println(v)
		}
	}
}

func UpdateRule() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = rule.UpdateRule(15, ipGeo.AllContinent, ipGeo.AllCountry, dns_common.DayEnd, dns_common.DayEnd, "www.google333.com", 100, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("update success")
	}
}

func DeleteRule() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = rule.Delete(15, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("update success")
	}
}

func main() {
	//AddRuleByRecordName()
	//AddRuleByRecordId()

	//QueryRuleByRecordName()
	//QueryRuleByRecordId()
	//
	//UpdateRule()
	//
	DeleteRule()
}
