package main

import (
	"fmt"
	"github_action/libre"
	"time"
)

func main() {
	fmt.Println("started")
	libre.LibreOfficeConvert("csv", "./test.xls")
	fmt.Println("done")
	time.Sleep(time.Minute * 25)
}

func ReturnOne() int {
	return 1
}
