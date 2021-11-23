package main

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/apex/log"
)

var (
	// This is seemingly required to prevent issues during concurrent uploads.
	libreOfficeMtx sync.Mutex
)

func main() {
	fmt.Println("started")
	LibreOfficeConvert("csv", "./test.xls")
	fmt.Println("done")
}

func ReturnOne() int {
	return 1
}

func LibreOfficeConvert(outputFormat string, fileName string) error {
	libreOfficeMtx.Lock()
	defer libreOfficeMtx.Unlock()
	cmd := exec.Command("soffice", "convert", "--headless", "--convert-to", outputFormat, "--outdir", "/tmp", fileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("converting file: %w", err)
	}

	log.WithField("output", string(output)).Info("executed libreoffice convert command")
	return nil
}
