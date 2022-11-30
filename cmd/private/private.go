package private

import (
	"fmt"

	"github.com/gosuri/uiprogress"
	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

func Scanner(data []string, filename string) string {
	utils.CreateFile(filename)
	uiprogress.Start()                  // start rendering
	bar := uiprogress.AddBar(len(data)) // Add a new bar

	bar.AppendCompleted()

	for _, ip := range data {
		utils.PingFunc(ip, filename)
		bar.Incr()
	}

	return fmt.Sprintf("Saved data to %s", filename)
}
