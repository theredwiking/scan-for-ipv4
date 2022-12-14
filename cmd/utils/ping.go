package utils

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
)

func PingFunc(ipAddr string, filename string) {
	pinger, err := ping.NewPinger(ipAddr)
	pinger.SetPrivileged(false)

	if err != nil {
		panic(err)
	}

	pinger.Count = 3

	pinger.OnFinish = func(stats *ping.Statistics) {
		if stats.PacketsRecv == 3 {
			data := JsonIp(ipAddr, "success")
			Save(string(data)+",", filename)
		} else {
			data := JsonIp(ipAddr, "failed")
			Save(string(data)+",", filename)
		}
	}

	pinger.Timeout = time.Second * 5

	err = pinger.Run()

	if err != nil {
		fmt.Println("Failed to ping target host:", err)
	}
}
