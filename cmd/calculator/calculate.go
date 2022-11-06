package calculator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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
		return
	}

	calculate(ip, subnet)
}

func calculate(baseip string, subnet string) {
	params := url.Values{}

	params.Add("cclass", "any")
	params.Add("csubnet", subnet)
	params.Add("cip", baseip)
	params.Add("ctype", "ipv4")
	params.Add("x", "0")
	params.Add("y", "0")

	resp, err := http.Get("https://www.calculator.net/ip-subnet-calculator.html?" + params.Encode())

	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body) // Log the request body

	if err != nil {
		fmt.Println(err)
		return
	}

	bodyString := string(body)
	fmt.Println(bodyString)
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
			return "", "", err
		}

		info := strings.Split(addr[0].String(), "/")

		return info[0], info[1], nil
	}
	return "", "", errors.New("are you even connected to a network?")
}
