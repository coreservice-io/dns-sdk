# dns-sdk

## How to use
```go
package main

import (
	"log"

	dns_common "github.com/coreservice-io/dns-common"
	dns_client "github.com/coreservice-io/dns-sdk"
	"github.com/coreservice-io/dns-sdk/record"
)

var token = "cxwinggdadlzdhpcyktmikjj" //token from dns-server
var domain = "mesoncdn.com" // managed domain
var endPoint = "http://127.0.0.1:9001" // request end point

//example Add a new dns record
func AddRecord() {
	// new a client
	client, err := dns_client.New(token, endPoint)
	if err != nil {
		log.Fatalln(err)
	}

	// add a new record under "mesoncdn.com"
	newRecord, err := record.Add(domain, "pullzone2", dns_common.TypeCNAME, 600, client)
	if err != nil {
		log.Println(err)
	}
	log.Println("newRecord:", newRecord)
}

func main() {
	AddRecord()
}
```

more 