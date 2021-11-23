package libre

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

func LibreOfficeConvert(outputFormat string, fileName string) error {
	libreOfficeMtx.Lock()
	defer libreOfficeMtx.Unlock()
	cmd := exec.Command("libreoffice", "convert", "--headless", "--convert-to", outputFormat, "--outdir", "/tmp", fileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("converting file: %w", err)
	}

	log.WithField("output", string(output)).Info("executed libreoffice convert command")
	return nil
}
