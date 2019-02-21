package main

import (
	"./mock"
	real2 "./real"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this mock retriever"}
	r = real2.Retriever{"www.baidu.com"}
	fmt.Println(download(r))
}
