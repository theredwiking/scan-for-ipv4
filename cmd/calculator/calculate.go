package calculator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

func CalcPath(path string) (string, error) {
	ip, subnet, err := localIp()

	if err != nil {
		return "", err
	}

	data, err := calculate(ip, subnet)

	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s/calc.txt", path)

	utils.CreateFile(filename)

	utils.Save(data, filename)

	return fmt.Sprintf("Saved data to %s", filename), nil
}

func CalcFile() (string, error) {
	ip, subnet, err := localIp()

	if err != nil {
		return "", err
	}

	data, err := calculate(ip, subnet)

	if err != nil {
		return "", err
	}

	utils.CreateFile("./calc.txt")

	utils.Save(data, "./calc.txt")

	return "Saved data to calc.txt", nil
}

func Calc() (string, error) {
	ip, subnet, err := localIp()

	if err != nil {
		return "", err
	}

	data, err := calculate(ip, subnet)

	if err != nil {
		return "", err
	}

	return data, nil
}

func calculate(baseip string, subnet string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://www.calculator.net/ip-subnet-calculator.html", nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:107.0) Gecko/20100101 Firefox/107.0")

	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("cclass", "any")
	q.Add("csubnet", subnet)
	q.Add("cip", baseip)
	q.Add("ctype", "ipv4")
	q.Add("x", "0")
	q.Add("y", "0")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var rgx = regexp.MustCompile(`<(td)>(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3} - \d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})</td>`)

	rs := rgx.FindStringSubmatch(string(body))

	return rs[2], nil
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
