package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	// change reader if encoding error occured
	// e := determineEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(resp.Body)
}

//func determineEncoding(
//	r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		return unicode.UTF8
//	}
//	e, _, _, := charset.DetermineEncoding(bytes, "")
//  return e
//}