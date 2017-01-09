package iprun

import (
	"time"
	"sync"
	_ "fmt"
)

//This is to be able to replaced by mock func in testing or imported
type DoSomething func(ip ...string) (string, error)

type Targets struct {
	Hosts []string
	Community string
	Oid string
	Action DoSomething
}


//i.e. GoSNMP(ip string, community string, oid string) (result map[string]string, err error)
//Input: ip address first and then whatever string you need
/*func doSomething (ip string) (map[string]string, error){
	//doNothing, just return map with ip address as key and empty string as value
	return map[string]string{ip:""}, nil
}*/
/* do Something with each string (IP address) in Target hosts within given timeout
Input: Action and slice of target strings (IP addresses)
Output: slice of IP addresses for which Action succeeded (no error)
 */
func RunIPs(p *Targets, timeout time.Duration) (map[string]string, error){
	concur := 100
	if len(p.Hosts) < concur {
		concur = len(p.Hosts)
	}
	q := make(chan struct{}, concur) // make 100 parallel connections
	wg := sync.WaitGroup{}
	result := make(map[string]string,len(p.Hosts))

	c := func(host ...string) {
		output , err := p.Action(host...)
		if err == nil {
			//fmt.Println(host,output)
			result[host[0]] = output
		}
		<-q
		wg.Done()
	}

	wg.Add(len(p.Hosts))
	for _, host := range p.Hosts {
		q <- struct{}{}
		go c(host,p.Community, p.Oid)
	}
	wg.Wait()
	return result, nil
}
