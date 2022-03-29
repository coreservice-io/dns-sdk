package main

import (
	"log"

	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/record"
)

var token = "cxwinggdadlzdhpcyktmikjj"
var domain = "mesoncdn.com"
var endPoint = "http://127.0.0.1:9001"

func AddRecord() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	newRecord, err := record.Add("pullzone1", "CNAME", 600, client)
	if err != nil {
		log.Println(err)
	}
	log.Println("newRecord:", newRecord)
	//recordName pullzone1
	//recordId
}

func ForbiddenRecordByName() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.ForbiddenByRecordName("pullzone1", true, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record forbidden")
	}
}

func ForbiddenRecordById() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.ForbiddenByRecordId(5, true, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record forbidden")
	}
}

func ActiveRecordByName() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.ForbiddenByRecordName("pullzone1", false, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record active")
	}
}

func ActiveRecordById() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.ForbiddenByRecordId(5, false, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record active")
	}
}

func QueryByNamePattern() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	records, totalCount, err := record.QueryByNamePattern("", 1, 0, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("records total count:", totalCount)
		for _, v := range records {
			log.Println(v)
		}
	}
}

func QueryByNameArray() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	records, err := record.QueryByGivenList([]string{"pullzone1"}, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("records:")
		for _, v := range records {
			log.Println(v)
		}
	}
}

func DeleteByName() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.DeleteByRecordName("pullzone1", "CNAME", client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("delete success")
	}
}

func DeleteById() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.DeleteByRecordId(5, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("delete success")
	}
}

func main() {
	//AddRecord()
	//
	//QueryByNameArray()
	QueryByNamePattern()

	//ForbiddenRecordByName()
	//ForbiddenRecordById()
	//
	//ActiveRecordByName()
	//ActiveRecordById()
	//
	//DeleteByName()
	//DeleteById()
}
