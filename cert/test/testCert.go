package main

import (
	"log"

	dns_client "github.com/coreservice-io/dns-client"
	"github.com/coreservice-io/dns-client/cert"
)

var token = "cxwinggdadlzdhpcyktmikjj"
var domain = "mesoncdn.com"
var endPoint = "http://127.0.0.1:9001"

func ApplyCert() {
	client, err := dns_client.New(token, domain, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	applyDomain := "www.somedomain.com"
	pullZoneName := "pullzonexxx"
	//before apply must set dns record in customer's dns server
	//1. CNAME  www.somedomain.com => pullzonexxx.mesoncdn.com
	//2. CNAME  _acme-challenge.example.somedomain.com => _acme-challenge.example.pullzonexxx.mesoncdn.com

	cert, key, err := cert.Apply(applyDomain, pullZoneName, client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("cert:", cert)
		log.Println("key:", key)
	}

}

func main() {
	//todo test in real server
	//ApplyCert()
}
