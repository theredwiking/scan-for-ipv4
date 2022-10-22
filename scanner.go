package main

import (
	"fmt"
	"os"
)

func main() {
	checkFile()

	first, second, third, fourth := 1, 1, 1, 0

	for true {

		for fourth < 255 {
			fourth++
			ipAddr := fmt.Sprintf("%d.%d.%d.%d", first, second, third, fourth)

			pingFunc(ipAddr)
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
