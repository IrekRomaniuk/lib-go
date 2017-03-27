package fetcher

import (
"testing"
"fmt"
)

func mock_get_page(url string) ([]byte, error){
	// mock your 'get_page()' function here
	fmt.Printf("Fake fetcher of %s\n", url)
	return []byte{}, nil
}

func TestFetcher(t *testing.T) {
	//--------Method1------------------------------------------------------------------
	fmt.Println("---------------------Method1---------------------")
	downloader(mock_get_page, "http://google.com")
	//--------Method2------------------------------------------------------------------
	fmt.Println("---------------------Method2---------------------")
	d := NewDownloader(mock_get_page)
	d.download("http://google.com")
	fmt.Println("-------------------------------------------------")

}
