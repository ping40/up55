package analyse

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"

	"github.com/axgle/mahonia"

	"github.com/ping40/up55/util"
)

func Anaylse(code string) {
	dataDir := util.MakeDataDirectory()
	var zcfzbResult [][]string
	var err error

	fileName := util.MakeZCFZBFileName(code)
	if err, zcfzbResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var lrbResult [][]string
	fileName = util.MakeLRBFileName(code)
	if err, lrbResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var xjllbResult [][]string
	fileName = util.MakeXJLLBFileName(code)
	if err, xjllbResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var growthResult [][]string
	fileName = util.MakeGrowthFileName(code)
	if err, growthResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var profitAbilityResult [][]string
	fileName = util.MakeProfitabilityFileName(code)
	if err, profitAbilityResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var repayingAbilityResult [][]string
	fileName = util.MakeRepayingAbilityFileName(code)
	if err, repayingAbilityResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	var operationAbilityResult [][]string
	fileName = util.MakeOperationAbilityFileName(code)
	if err, operationAbilityResult = csvRead(filepath.Join(dataDir, fileName)); err != nil {
		panic(err)
	}

	genHumanRead(code, zcfzbResult, lrbResult, xjllbResult, growthResult, profitAbilityResult, repayingAbilityResult, operationAbilityResult)
}

func csvRead(fullFileName string) (error, [][]string) {
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")

	// Open the file
	csvfile, err := os.Open(fullFileName)
	if err != nil {
		return err, nil
	}
	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)

	r.FieldsPerRecord = -1

	result := make([][]string, 0)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err, nil
		}

		row := make([]string, 0)

		for i, v := range record {
			if i == len(record)-1 && v == "" {
				break
			}
			if i == 0 {
				v = enc.ConvertString(v)
			}
			row = append(row, v)

		}

		result = append(result, row)
	}

	return nil, result
}
