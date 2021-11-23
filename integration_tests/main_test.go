// +build integration

package integrationtests_tests

import (
	"encoding/csv"
	"github_action/libre"
	"os"
	"testing"

	"github.com/apex/log"
)

func TestLibreOfficeConvert(t *testing.T) {
	libre.LibreOfficeConvert("csv", "../files/test.xls")
	// defer os.Remove("/tmp/test.csv")
	f, err := os.Open("/tmp/test.csv")
	if err != nil {
		log.WithError(err).Error("failed to open tmp csv file")
		t.Errorf("failed to open file")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		t.Errorf("failed to read created csv file")
	}

	o, err := os.Open("../files/test.csv")
	if err != nil {
		t.Errorf("failed to open existing csv file: %s", err.Error())
	}

	csvReaderO := csv.NewReader(o)
	recordsO, err := csvReaderO.ReadAll()
	if err != nil {
		t.Errorf("failed to read existing csv file: %s", err.Error())
	}
	if recordsO[1][0] != records[0][0] {
		t.Errorf("files are not equal")
	}
}
