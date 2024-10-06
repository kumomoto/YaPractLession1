package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"yapractlession1/cmd"
)

func main() {

	var errCount int

	for {

		resp, respErr := http.Get("http://srv.msk01.gigacorp.local/_stats")
		if respErr != nil {
			fmt.Println(respErr)
		}

		if resp.StatusCode != 200 {
			if errCount > 3 {
				fmt.Println("Unable to fetch server statistic")
				continue
			} else {
				errCount += 1
				continue
			}
		} else {
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				fmt.Println(err)
			}

			metrics, bodyErr := parseBody(body)

			if bodyErr != nil {
				if errCount > 3 {
					fmt.Println("Unable to fetch server statistic")
					continue
				} else {
					errCount += 1
					continue
				}
			} else {
				cmd.GetLaDesicions(metrics[0])
				cmd.GetRAMDesicions(metrics[1], metrics[2])
				cmd.GetDiskDesicions(metrics[3], metrics[4])
				cmd.GetNetworkDesicions(metrics[5], metrics[6])
			}

		}

		time.Sleep(5 * time.Second)

	}

}

func parseBody(body []byte) ([]int, error) {
	var err error = nil

	bodyString := string(body)

	bodyString = strings.Trim(bodyString, "\n")

	bodySlice := strings.Split(bodyString, ",")

	values := []int{}

	for _, body := range bodySlice {
		item, convertErr := strconv.Atoi(strings.Trim(body, " "))

		values = append(values, item)

		if convertErr != nil {
			err = convertErr
			break
		}
	}
	return values, err
}
