package fetcher

import (
	"testing"
	"fmt"
)

type fakeFetcher struct{}

func (fetcher fakeFetcher) Fetch(url string) ([]byte, error) {
	fmt.Printf("Fake fetcher%s\n", url)
	return []byte{}, nil
}

func TestFetcher(t *testing.T) {
	var fake fakeFetcher
	//var real realFetcher
	Crawl(fake)
	//Crawl(real)
}
