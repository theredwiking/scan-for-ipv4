package main

import (
	"fmt"

	"github.com/go-ping/ping"
)

func pingFunc(ipAddr string) {
	pinger, err := ping.NewPinger(ipAddr)

	if err != nil {
		panic(err)
	}

	pinger.Count = 3

	err = pinger.Run()

	if err != nil {
		panic(err)
	}

	stats := pinger.Statistics()

	if stats.PacketLoss > 0 || stats.PacketsRecv < 3 {
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n", stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		save(ipAddr, "failed")
	} else {
		save(ipAddr, "success")
	}
}
