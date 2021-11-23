package main

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/apex/log"
)

func TestReturnOne(t *testing.T) {
	res := ReturnOne()
	if res != 1 {
		t.Errorf("failed to return 1: %d", res)
	} else {
		t.Logf("test succeeded")
	}
}

func TestLibreOfficeConvert(t *testing.T) {
	LibreOfficeConvert("csv", "test.xls")
	defer os.Remove("/tmp/test.csv")
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

	o, err := os.Open("./test.csv")
	if err != nil {
		t.Errorf("failed to open existing csv file")
	}

	csvReaderO := csv.NewReader(o)
	recordsO, err := csvReaderO.ReadAll()
	if err != nil {
		t.Errorf("failed to read existing csv file")
	}
	if recordsO[0][0] != records[0][0] {
		t.Errorf("files are not equal")
	}
}
