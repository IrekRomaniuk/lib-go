package gosnmp

import (
	"testing"
)

func TestGoSNMP(t *testing.T) {

	type input struct {
		ip, community, oid string
	}
	/*$ snmpget -v1 -c public 10.199.107.1 system.sysContact.0
	$ snmpget -v2c -c public 10.199.107.1 .1.3.6.1.2.1.1.4.0
	SNMPv2-MIB::sysContact.0 = STRING: irek romaniuk*/
	var inputTest = input{"10.199.107.1","public","1.3.6.1.2.1.1.4.0"}  //system.sysContact.0
	want := "irek romaniuk"

//func GoSNMP(ip string, community string, oid string) (result string, err error)
	if output, _ := GoSNMP(inputTest.ip, inputTest.community, inputTest.oid); output != want {
		t.Fatalf("sysContact: got %v, want %v",output, want)
	}
}
