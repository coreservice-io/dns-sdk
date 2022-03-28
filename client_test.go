package dns_client

import (
	"log"
	"testing"
)

func Test_NewClient(t *testing.T) {
	client, err := New("ctijtcoupixkasnjihxfehgt", "coreservice.io", "http://127.0.0.1:9001")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(client)
}
