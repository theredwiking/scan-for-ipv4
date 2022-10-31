package calculator

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

func CalcPath(path string) {
	// ip, subnet, err := localIp()

	// if err != nil {
	// 	fmt.Println(err)
	// }
	utils.CreateFile(fmt.Sprintf("%s/calc.txt", path))
}

func CalcFile() {
	// ip, subnet, err := localIp()

	// if err != nil {
	// 	fmt.Println(err)
	// }
	utils.CreateFile("./calc.txt")
}

func Calc() {
	ip, subnet, err := localIp()

	if err != nil {
		fmt.Println(err)
	}

	calculate(ip, subnet)
}

func calculate(baseip string, subnet string) {
	
}

func localIp() (string, string, error) {
	ifaces, err := net.Interfaces()

	if err != nil {
		return "", "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 || iface.Name == "virbr0" {
			continue // interface down & loopback interface
		}

		addr, err := iface.Addrs()

		if err != nil {
			fmt.Println(err)
		}

		info := strings.Split(fmt.Sprintf("%s", addr[0]), "/")

		return info[0], info[1], nil
	}
	return "", "", errors.New("Are you even connected to a network?")
}
