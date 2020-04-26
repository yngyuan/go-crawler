package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main()  {
	fmt.Println("hello crawler")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}
	// change reader if encoding error occured
	// e := determineEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	printCityList(all)
}

func printCityList(contents []byte)  {
	re := regexp.MustCompile(`<a href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	maches := re.FindAllSubmatch(contents, -1)
	for _, m := range maches{
		fmt.Printf("City: %s, URL: %s\n", m[2], m[1] )
	}
	fmt.Printf("Matches Found %d\n", len(maches))
}

//func determineEncoding(
//	r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _, := charset.DetermineEncoding(bytes, "")
//}

