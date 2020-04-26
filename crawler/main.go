package main

import (
	"fmt"
	"net/http"
)

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

func main()  {
	fmt.Println("hello crawler")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
