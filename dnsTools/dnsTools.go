package dnsTools

import (
	"net"
)

func CheckCNAME(domain string) (dest string, err error) {
	dest, err = net.LookupCNAME(domain)
	return
}
