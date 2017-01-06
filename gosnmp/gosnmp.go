package gosnmp

import (
	"github.com/alouca/gosnmp"
	"log"
)

func GoSNMP(ip string, community string, oid string) (result string, err error) {
	s, err1 := gosnmp.NewGoSNMP(ip, community, gosnmp.Version2c, 5)
	if err1 != nil {
		return "", err
	}
	resp, err2 := s.Get(oid)
	if err2 == nil {
		for _, v := range resp.Variables {
			switch v.Type {
			case gosnmp.OctetString:
				log.Printf("Response: %s : %s : %s \n", v.Name, v.Value.(string), v.Type.String())
				result = v.Value.(string)
			}
		}
	} else {
		return "", err2
	}
	return result, nil
}

