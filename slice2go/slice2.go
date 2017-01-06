package slice2go

import (
	"time"
	"sync"
)

type DoSomething func(ip string) (string, error)

type Target struct {
	hosts []string
	Action DoSomething
}

func doSomething (ip ...string) (string, error){ //ip address first and then whatever string you need
	return ip[0], nil  //doNothing
}
/* do Something with each string (IP address) in Target hosts within given timeout
Input: Action and slice of target strings (IP addresses)
Output: slice of IP addresses for which Action succeeded (no error)
 */
func Slice2go(p *Target, timeout time.Duration) ([]string, error){
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
			result=append(result, host)
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
