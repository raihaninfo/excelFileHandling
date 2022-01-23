package main

import (
	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet1")
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Amount")
	f.SetCellValue("Sheet1", "A2", "Abdur Rahman")
	f.SetCellValue("Sheet1", "B2", 33)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	err := f.SaveAs("file1.xlsx")
	if err != nil {
		panic(err)
	}
}
