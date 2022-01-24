package main

import (
    "fmt"

    "github.com/xuri/excelize/v2"
)

func main() {
    f, err := excelize.OpenFile("file1.xlsx")
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
    }
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}
