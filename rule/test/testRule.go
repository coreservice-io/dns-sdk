package main

import (
	"log"

	dns_common "github.com/coreservice-io/dns-common"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/rule"
	"github.com/coreservice-io/ipGeo"
)

var token = "cxwinggdadlzdhpcyktmikjj"
var domainName = "mesoncdn.com"
var endPoint = "http://127.0.0.1:9001"

func AddRule() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	newRules := []*rule.NewRuleData{
		{ipGeo.AllContinent, ipGeo.AllCountry, "www.google.com", 100},
		{ipGeo.AllContinent, ipGeo.AllCountry, "www.youtube.com", 100},
	}
	newRule, err := rule.AddRule(domainName, "pullzone1", dns_common.TypeCNAME, newRules, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("newRule", newRule)
	}
}

func DeleteRule() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = rule.DeleteRules(domainName, "pullzone1", dns_common.TypeCNAME, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("update success")
	}
}

func main() {
	AddRule()
	DeleteRule()
}
