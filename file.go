package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Ip struct {
	IpAdrr, Status string
}

func save(ipAddr string, status string) {
	f, err := os.OpenFile("result.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	data := Ip{
		IpAdrr: ipAddr,
		Status: status,
	}

	entry, _ := json.Marshal(data)

	_, err = fmt.Fprintln(f, string(entry))

	if err != nil {
		fmt.Println(err)
	}
}

func checkFile() {
	if _, err := os.Stat("result.json"); err == nil {
		os.Remove("result.json")
	}

	f, err := os.Create("result.json")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
}
