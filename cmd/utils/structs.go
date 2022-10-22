package utils

import (
	"encoding/json"
)

type Ip struct {
	IpAdrr, Status string
}

func JsonData(ipAddr string, status string) []byte {

	data := Ip{
		IpAdrr: ipAddr,
		Status: status,
	}

	obj, _ := json.Marshal(data)
	return obj
}
