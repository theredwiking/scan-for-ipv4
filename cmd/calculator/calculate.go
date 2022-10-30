package calculator

import (
	"fmt"

	"github.com/theredwiking/scan-for-ipv4/cmd/utils"
)

func CalcPath(path string) {
	utils.CreateFile(fmt.Sprintf("%s/calc.json", path))
}

func CalcFile() {
	utils.CreateFile("./calc.json")
}

func Calc() {
	fmt.Println("Calculating")
}
