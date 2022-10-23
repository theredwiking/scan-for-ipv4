package public

import (
	"fmt"
	"sync"

	"github.com/gosuri/uiprogress"
	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

var wg sync.WaitGroup

func Scanner(filename string) {
	utils.CreateFile(filename)

	wg.Add(3)
	go ping(1, filename)
	go ping(85, filename)
	go ping(170, filename)

	wg.Wait()
}

func ping(first int, filename string) {
	uiprogress.Start()            // start rendering
	bar := uiprogress.AddBar(100) // Add a new bar

	second, third, fourth := 0, 0, 0

	firstLimit := first + 85

	bar.AppendCompleted()
	bar.PrependElapsed()

	for true {

		for fourth < 255 {
			fourth++
			ipAddr := fmt.Sprintf("%d.%d.%d.%d", first, second, third, fourth)

			utils.PingFunc(ipAddr, filename)
		}
		bar.Incr()
		fourth = 0

		if third != 255 {
			third++
		}
		if second != 255 && third == 255 {
			third = 0
			second++

			switch second {
			case 64:
				if first == 100 {
					second++
				}
			case 254:
				if first == 169 {
					second++
				}
			case 16:
				if first == 172 {
					second += 16
				}
			case 0:
				if first == 192 && third == 0 || first == 192 && third == 2 || first == 203 && third == 113 {
					third++
				}
			case 31:
				if first == 192 && third == 196 {
					third++
				}
			case 52:
				if first == 192 && third == 193 {
					third++
				}
			case 88:
				if first == 192 && third == 99 {
					third++
				}
			case 168:
				if first == 192 {
					second++
				}
			case 175:
				if first == 192 && third == 48 {
					third++
				}
			case 18:
				if first == 198 && third == 0 {
					second += 2
				}
			case 51:
				if first == 198 && third == 100 {
					third++
				}
			}
		}
		if first != firstLimit && second == 255 {
			third = 0
			second = 0
			first++

			switch first {
			case 10, 127:
				first++
			case 240:
				first += 15
			}
		}
		if first == firstLimit {
			wg.Done()
		}
	}
}
