package main

import (
	"io"
	"net/http"
	"strings"
	"time"
	"yapractlession1/cmd"
)

var errCount int

func main() {

	var resp *http.Response

	for {
		resp, respErr := http.Get("srv.msk01.gigacorp.local/_stats")

		if respErr != nil {
			println(respErr)
		}

		ok := CheckContentType(resp)

		if !ok {
			if errCount >= 3 {
				println("Unable to fetch server statistic")
			} else {
				errCount = errCount + 1
				time.Sleep(time.Second * 60)
				continue
			}
		}

		if resp.StatusCode != 200 {
			if errCount >= 3 {
				println("Unable to fetch server statistic")
			} else {
				errCount = errCount + 1
				time.Sleep(time.Second * 60)
				continue
			}
		}

	}

	body, bodyErr := io.ReadAll(resp.Body)
	resp.Body.Close()

	if bodyErr != nil {
		println(bodyErr)
	}

	currentSysProp := cmd.InitSystemProp(body)

	println(currentSysProp)

}

func CheckContentType(resp *http.Response) bool {

	var answer bool

	contentType := resp.Header.Get("Content-Type")

	if strings.Compare(contentType, "text/html; charset=utf-8") == 0 {
		answer = true
	} else {
		answer = false
	}

	return answer
}
