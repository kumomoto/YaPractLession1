package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dariubs/percent"
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
				GetLaDesicions(metrics[0])
				GetRAMDesicions(metrics[1], metrics[2])
				GetDiskDesicions(metrics[3], metrics[4])
				GetNetworkDesicions(metrics[5], metrics[6])
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

func GetNetworkDesicions(netBandwith, netLoad int) {
	var ans string = ""

	if percent.PercentOf(netLoad, netBandwith) > 90 {
		ans = fmt.Sprintf("Network bandwidth usage high: %v Mbit/s available", math.Trunc(((float64(netBandwith - netLoad)) / 125000 / 8)))
		fmt.Println(ans)
	}
}

func GetDiskDesicions(diskTotal, diskResident int) {
	var ans string = ""

	if percent.PercentOf(diskResident, diskTotal) > 90 {
		ans = fmt.Sprintf("Free disk space is too low: %.f Mb left", math.Trunc((float64(diskTotal-diskResident))/(1024*1024)))
		fmt.Println(ans)
	}
}

func GetRAMDesicions(ram, resram int) {

	var ans string = ""

	perc := percent.PercentOf(resram, ram)

	result := fmt.Sprintf("%.f", math.Trunc(perc))

	if perc > 80 {
		ans = fmt.Sprintf("Memory usage too high: " + result + "%%")
		fmt.Println(ans)
	}
}

func GetLaDesicions(la int) {
	var ans string = ""

	if la > 30 {
		ans = fmt.Sprintf("Load Average is too high: %d", la)
		fmt.Println(ans)
	}
}
