package main

import (
	"log"

	dns_common "github.com/coreservice-io/dns-common"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/record"
)

var token = "cxwinggdadlzdhpcyktmikjj"
var domainName = "mesoncdn.com"
var endPoint = "http://127.0.0.1:9001"

func AddRecord() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	newRecord, err := record.Add(domainName, "pullzone2", dns_common.TypeCNAME, 600, client)
	if err != nil {
		log.Println(err)
	}
	log.Println("newRecord:", newRecord)
}

func ForbiddenRecord() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	forbidden := true
	err = record.Update(domainName, "pullzone1", dns_common.TypeCNAME, &forbidden, nil, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record forbidden")
	}
}

func ActiveRecord() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	forbidden := false
	err = record.Update(domainName, "pullzone1", dns_common.TypeCNAME, &forbidden, nil, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record forbidden")
	}
}

func QueryRecords() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	records, totalCount, err := record.Query(domainName, []string{"pullzone1"}, dns_common.TypeCNAME, 0, 0, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("records total count:", totalCount)
		for _, v := range records {
			log.Println(v)
		}
	}
}

func DeleteRecord() {
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.Delete(domainName, "test", dns_common.TypeCNAME, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("delete success")
	}
}

func main() {
	AddRecord()

	ForbiddenRecord()
	ActiveRecord()

	QueryRecords()

	DeleteRecord()
}
