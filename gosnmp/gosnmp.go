package gosnmp

import (
	"github.com/alouca/gosnmp"
	_ "log"
)
/*
Input: "10.199.107.1","public","1.3.6.1.2.1.1.4.0"
Output: "irek romaniuk"
 */
func GoSNMP(input ...string) (result string, err error) {
	ip := input[0]
	community := input[1]
	oid := input[2]
	s, err1 := gosnmp.NewGoSNMP(ip, community, gosnmp.Version2c, 5000)
	if err1 != nil {
		return "", err1
	}
	resp, err2 := s.Get(oid)
	if err2 == nil {
		for _, v := range resp.Variables {
			switch v.Type {
			case gosnmp.OctetString:
				//log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
				result = v.Value.(string)
			}
		}
	} else {
		return "", err2
	}
	return result, nil
}

