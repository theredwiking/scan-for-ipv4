package utils

import (
	"fmt"
	"os"
)

func Save(data string) {
	f, err := os.OpenFile("result.json", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fprintln(f, data)

	if err != nil {
		fmt.Println(err)
	}
}

func CheckFile() {
	if _, err := os.Stat("result.json"); err == nil {
		os.Remove("result.json")
	}

	f, err := os.Create("result.json")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
}
