package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	Txt string
}

func (r Retriever) Get(url string) string {
	//
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(result)
}
