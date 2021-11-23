package main

import (
	"fmt"
	"github_action/libre"
)

func main() {
	fmt.Println("started")
	libre.LibreOfficeConvert("csv", "./test.xls")
	fmt.Println("done")
}

func ReturnOne() int {
	return 1
}
