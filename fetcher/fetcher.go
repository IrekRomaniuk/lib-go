package fetcher
//http://stackoverflow.com/questions/19167970/mock-functions-in-go
import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"fmt"
)

//--------Method1------------------------------------------------------------------

type PageGetter func(url string) ([]byte, error)

func downloader(pageGetterFunc PageGetter, url string) error{
	content, err := pageGetterFunc(url)
	if err != nil {
		return err
	}
	fmt.Printf("Crawling..%d\n", len(content))

	return nil
}


func get_page(url string) ([]byte, error) {
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

//--------Method2------------------------------------------------------------------
type Downloader struct {
	get_page PageGetter
}

func NewDownloader(pg PageGetter) *Downloader {
	return &Downloader{get_page: pg}
}

//fetching url and printing content length
func (d *Downloader) download(url string) ([]byte,error) {

	content, err := d.get_page(url)
	if err != nil {
		return []byte{},err
	}
	fmt.Printf("Crawling..%d\n", len(content))

	return content, nil
}