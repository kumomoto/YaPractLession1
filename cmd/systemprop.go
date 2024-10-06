package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type systemProp struct {
	la           int
	ramTotal     int
	ramResident  int
	diskTotal    int
	diskResident int
	netBandwith  int
	netLoad      int
}

func InitSystemProp(body []byte) *systemProp {
	values := parseBody(body)

	prop := &systemProp{
		la:           values[0],
		ramTotal:     values[1],
		ramResident:  values[2],
		diskTotal:    values[3],
		diskResident: values[4],
		netBandwith:  values[5],
		netLoad:      values[6],
	}

	return prop
}

func (s *systemProp) GetEvaluete() string {
	var ans string

	desicionsMap := make(map[string]string)

	desicionsMap["la"] = getLaDesicions(s.la)
	desicionsMap["ram"] = getRamDesicions(s.ramTotal, s.ramResident)
	desicionsMap["disk"] = getDiskDesicions(s.diskTotal, s.diskResident)

	for _, v := range desicionsMap {
		if len(v) != 0 {
			ans = ans + v + "\n"
		}
	}

	return ans
}

func getNetworkDesicions(netBandwith, netLoad int) string {
	var ans string

	if (netLoad/netBandwith)*100 > 90 {
		ans = fmt.Sprintf("Network bandwidth usage high: %v Mbit/s available", ((netBandwith - netLoad) / (1024 * 1024)))
	}

	return ans
}

func getDiskDesicions(diskTotal, diskResident int) string {
	var ans string

	if (diskResident/diskTotal)*100 > 90 {
		ans = fmt.Sprintf("Free disk space is too low: %v Mb left", (diskTotal-diskResident)/(1024*1024))
	}

	return ans
}

func getRamDesicions(ram, resram int) string {

	var ans string

	if (resram/ram)*100 > 80 {
		ans = fmt.Sprintf("Memory usage too high: %v", (resram/ram)*100)
	}

	return ans
}

func getLaDesicions(la int) string {
	var ans string

	if la > 30 {
		ans = fmt.Sprintf("Load Average is too high: %v", la)
	}

	return ans
}

func parseBody(body []byte) []int {
	bodyString := string(body)

	bodySlice := strings.Split(bodyString, ",")

	values := make([]int, len(bodySlice))

	for i := 0; i < len(bodySlice); i++ {
		item, err := strconv.Atoi(bodySlice[i])

		values[i] = item

		if err != nil {
			println(err)
		}
	}
	return values
}
