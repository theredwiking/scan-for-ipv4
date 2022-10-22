package public

import (
	"fmt"
	"os"

	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

func Scanner() {
	utils.CheckFile()

	first, second, third, fourth := 1, 1, 1, 0

	for true {

		for fourth < 255 {
			fourth++
			ipAddr := fmt.Sprintf("%d.%d.%d.%d", first, second, third, fourth)

			utils.PingFunc(ipAddr)
		}

		fourth = 0

		if third != 255 {
			third++
		}
		if second != 255 && third == 255 {
			third = 0
			second++
		}
		if first != 255 && second == 255 {
			third = 0
			second = 0
		}
		if first == 255 {
			os.Exit(0)
		}
	}
}
