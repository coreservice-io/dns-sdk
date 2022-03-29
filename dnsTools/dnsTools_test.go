package dnsTools

import (
	"log"
	"testing"
)

func Test_CNAME(t *testing.T) {
	log.Println(CheckCNAME("testcname.mesoncdntest.live"))
}
