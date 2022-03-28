package main

import (
	"log"

	dns_client "github.com/coreservice-io/dns-client"
	"github.com/coreservice-io/dns-client/record"
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
}

func ForbiddenRecord() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.Forbidden("pullzone1", true, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record forbidden")
	}
}

func ActiveRecord() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.Forbidden("pullzone1", false, client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("record active")
	}
}

func Query() {
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

func Delete() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	err = record.Delete("pullzone1", client)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("delete success")
	}
}

func main() {
	//AddRecord()
	//ForbiddenRecord()
	//ActiveRecord()
	//Query()
	//Delete()
}
