package cmd

import (
	"fmt"
	"math"

	"github.com/dariubs/percent"
)

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

	result := fmt.Sprintf("%v", math.RoundToEven(perc))

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
