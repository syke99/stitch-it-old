package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func generateExcelPattern(patternNm string, colArr [][]string, colNum map[string]int) string {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "Thread Colors for this Pattern")
	f.MergeCell("Sheet1", "A1", "B2")

	style, _ := f.NewStyle(`{"alighnment":{"horizontal":"center","vertical":"center"},"font":{"bold":true,"italic":false}}`)

	f.SetCellStyle("Sheet1", "A1", "B2", style)

	for key, element := range colNum {

		var rowIndx int = 3

		rowStr := strconv.Itoa(rowIndx)

		var col1 = "A" + rowStr
		var col2 = "B" + rowStr

		threadNum := strconv.Itoa(element)

		f.SetCellValue("Sheet1", col1, threadNum)
		f.SetCellStyle("Sheet1", col1, col1, style)

		f.SetCellValue("Sheet1", col2, key)
		f.SetCellStyle("Sheet1", col2, col2, style)

		rowIndx++
	}

	indx := f.NewSheet("Sheet2")

	f.SetActiveSheet(indx)

	var row int = 1

	rowStr := strconv.Itoa(row)

	alph := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for _, i := range colArr {

		for _, j := range i {

			alphI := 0

			alphCntr := 1

			if alphCntr != 25 {

				col := (strings.Repeat(alph[alphI], alphCntr)) + rowStr

				style := []string{col, col, "2", "2", "2", "2"}

				s, _ := f.NewStyle(fmt.Sprintf(`{"border":[{"type":"left","color":"#000000","style":%s},{"type":"top","color":"#000000","style":%s},{"type":"bottom","color":"#000000","style":%s},{"type":"right","color":"#000000","style":%s}]}`, style[2], style[3], style[4], style[5]))

				f.SetCellValue("Sheet2", col, colNum[j])
				f.SetCellStyle("Sheet2", col, col, s)

				alphCntr++

			} else {
				alphI = 0
				alphCntr++

				col := (strings.Repeat(alph[alphI], alphCntr)) + rowStr

				style := []string{col, col, "2", "2", "2", "2"}

				s, _ := f.NewStyle(fmt.Sprintf(`{"border":[{"type":"left","color":"#000000","style":%s},{"type":"top","color":"#000000","style":%s},{"type":"bottom","color":"#000000","style":%s},{"type":"right","color":"#000000","style":%s}]}`, style[2], style[3], style[4], style[5]))

				f.SetCellValue("Sheet2", col, colNum[j])
				f.SetCellStyle("Sheet2", col, col, s)

				alphCntr++
			}
		}
		row++
	}

	fileName := patternNm + ".xlsx"

	if err := f.SaveAs(fileName); err != nil {
		genErr := "failed to generate and save excel pattern"

		return genErr
	}

	return ""
}
