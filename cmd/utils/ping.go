package utils

import (
	"github.com/go-ping/ping"
)

func PingFunc(ipAddr string) {
	pinger, err := ping.NewPinger(ipAddr)

	if err != nil {
		panic(err)
	}

	pinger.Count = 3

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketLoss > 0 || stats.PacketsRecv < 3 {
			data := JsonData(ipAddr, "failed")
			Save(string(data))
		} else {
			data := JsonData(ipAddr, "success")
			Save(string(data))
		}
	}

	err = pinger.Run()

	if err != nil {
		panic(err)
	}
}
