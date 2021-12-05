package main

import (
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func generateExcelPattern(patternNm string, colArr [][]string, colNum map[string]int) {
	f := excelize.NewFile()

	for num, _ := range colArr {
		if num == 0 {

			f.SetCellValue("Sheet1", "A1", "Thread Colors for this Pattern")
			f.MergeCell("Sheet1", "A1", "B2")

			style, _ := f.NewStyle(`{"alighnment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":false}}`)

			f.SetCellStyle("Sheet1", "A1", "B2", style)

			for key, element := range colNum {

				var rowIndx int = 3

				var rowStr = strconv.Itoa(rowIndx)

				var col1 = "A" + rowStr
				var col2 = "B" + rowStr

				var threadNum = strconv.Itoa(element)

				f.SetCellValue("Sheet1", col1, threadNum)
				f.SetCellStyle("Sheet1", col1, col1, style)

				f.SetCellValue("Sheet1", col2, key)
				f.SetCellStyle("Sheet1", col2, col2, style)

				rowIndx++
			}
		}
		// for _, j := range i {

		// }
	}

	fileName := patternNm + ".xlsx"

	if err := f.SaveAs(fileName); err != nil {
		log.Fatal(err)
	}
}
