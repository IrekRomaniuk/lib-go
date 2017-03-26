package fetcher
//http://relistan.com/writing-testable-apps-in-go/
import (
	"io/ioutil"
	"net/http"
	"fmt"
	"crypto/tls"
)

type HttpResponseFetcher interface {
	Fetch(url string) ([]byte, error)
}

type realFetcher struct{}

func Crawl(fetcher HttpResponseFetcher) error {
	_, err := fetcher.Fetch("http://google.com")
	if err == nil {
		//Do Smth
		fmt.Println("Crawling..")
		if err == nil {
			return nil
		}
	}
	return err
}


func (fetcher realFetcher) Fetch(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		return []byte{}, err
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	resp.Body.Close()
	return htmlData, nil
}

/*func (fetcher realFetcher) Fetch(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	fmt.Printf("Real fetcher %s\n", url)
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return contents, nil
}*/