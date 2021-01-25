package analyse

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	iconv "github.com/djimenez/iconv-go"

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

	fmt.Println("zcfzbResult[0]:  ", zcfzbResult[0])
	fmt.Println("lrbResult[0]:  ", lrbResult[0])
	fmt.Println("xjllbResult[0]:  ", xjllbResult[0])

	f, err := excelize.OpenFile(filepath.Join(dataDir, fileName))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func csvRead(fullFileName string) (error, [][]string) {
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

		fmt.Printf("Record has %d columns.\n", len(record))
		city, _ := iconv.ConvertString(record[0], "gb2312", "utf-8")

		fmt.Print(" city: ", city)
		row = append(row, city)
		for i, v := range record {
			if i > 0 {
				fmt.Printf("%s,", v)
				row = append(row, v)
			}
		}
		fmt.Println("")

		result = append(result, row)
	}

	return nil, result
}
