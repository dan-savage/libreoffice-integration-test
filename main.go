package main

import (
	"fmt"
	"github_action/libre"
	"time"
)

func main() {
	time.Sleep(time.Second * 25)
	libre.LibreOfficeConvert("csv", "./test.xls")
	fmt.Println("done")
	time.Sleep(time.Minute * 25)
}

func ReturnOne() int {
	return 1
}
