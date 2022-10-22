package main

import (
	"fmt"
)

func main() {

	checkFile()

	first := 0

	for first < 255 {
		first++
		second := 0

		for second < 255 {
			second++
			third := 0

			for third < 255 {
				third++
				fourth := 0

				for fourth < 255 {
					fourth++

					ipAddr := fmt.Sprintf("%d.%d.%d.%d", first, second, third, fourth)

					pingFunc(ipAddr)
				}
			}
		}
	}
}
