package iprun

import (
	"time"
	"sync"
)

//This is to be able to replaced by mock func in testing or imported
type DoSomething func(ip ...string) (map[string]string, error)

type Targets struct {
	hosts []string
	community string
	oid string
	Action DoSomething
}


//i.e. GoSNMP(ip string, community string, oid string) (result map[string]string, err error)
//Input: ip address first and then whatever string you need
func doSomething (ip string) (map[string]string, error){
	//doNothing, just return map with ip address as key and empty string as value
	return map[string]string{ip:""}, nil
}
/* do Something with each string (IP address) in Target hosts within given timeout
Input: Action and slice of target strings (IP addresses)
Output: slice of IP addresses for which Action succeeded (no error)
 */
func RunIPs(p *Targets, timeout time.Duration) ([]string, error){
	concur := 100
	if len(p.hosts) < concur {
		concur = len(p.hosts)
	}
	q := make(chan struct{}, concur) // make 100 parallel connections
	wg := sync.WaitGroup{}
	var result []string

	c := func(host string) {
		_ , err := p.Action(host)
		if err == nil {
			result=append(result,host)
		}
		<-q
		wg.Done()
	}

	wg.Add(len(p.hosts))
	for _, host := range p.hosts {
		q <- struct{}{}
		go c(host)
	}
	wg.Wait()
	return result, nil
}
