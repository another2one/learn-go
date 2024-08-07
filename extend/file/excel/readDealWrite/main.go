package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"learn-go/common/funcs"
	"log"
	"os"
	"strconv"
)

var (
	excelFile = funcs.ProjectPath + "extend/file/excel/readDealWrite/test.xlsx"
)

// https://xuri.me/excelize/zh-hans/cell.html
// 生成柱状图标
func main() {
	createTestFile()
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		log.Fatalf("read excel file [%s] err: %s", excelFile, err)
		return
	}

	sheet := "Sheet1"

	rows, err := f.GetRows(sheet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", rows)

	if err = f.SetCellValue(sheet, "E1", "count"); err != nil {
		log.Fatalf("write excel cell [E1] err: %s", err)
	}

	for row := 2; row <= 4; row++ {
		count := 0
		for _, col := range []byte("BCD") {
			cellStr := string(col) + strconv.Itoa(row)
			cellValue, err := f.GetCellValue(sheet, cellStr)
			if err != nil {
				log.Fatalf("read excel cell [%s] err: %s", cellValue, err)
			}
			valueInt, err := strconv.Atoi(cellValue)
			if err != nil {
				log.Fatalf("cell [%s] = %+v parse err: %s", cellValue, valueInt, err)
			}
			count += valueInt
		}
		cellStr := "E" + strconv.Itoa(row)
		if err = f.SetCellValue(sheet, cellStr, count); err != nil {
			log.Fatalf("set excel cell [%s] = %s err: %s", cellStr, count, err)
		}
	}

	// 根据指定路径保存文件
	if err := f.Save(); err != nil {
		fmt.Println(err)
	}
}

func createTestFile() {
	err := os.Remove(excelFile)
	if err != nil {
		fmt.Println(err)
	}
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	// 根据指定路径保存文件
	if err := f.SaveAs(excelFile); err != nil {
		fmt.Println(err)
	}
}
