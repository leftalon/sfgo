package fileutil

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
)

// 导出excel
func ExportExcel(data [][]interface{}, title []string, sheetName string, fileName string) (string, error) {
	// 判断filepath文件是否存在
	if IsFile(fileName) {
		return fileName, errors.New("文件已经存在")
	}
	f := xlsx.NewFile()
	// Create a new sheet.
	sheet, err := f.AddSheet(sheetName)
	if err != nil {
		return fileName, err
	}
	// 处理字段
	rowColumn := sheet.AddRow()
	for _, value := range title {
		cell := rowColumn.AddCell()
		cell.SetValue(value)
	}
	// 处理数据
	for _, v := range data {
		rowData := sheet.AddRow()
		for _, value := range v {
			cell := rowData.AddCell()
			cell.SetValue(value)

		}
	}
	err = f.Save(fileName)
	return fileName, err

}

func ExportCsv(data [][]interface{}, title []string, sheetName string, fileName string) (string, error) {
	var f *os.File
	var valueList [][]string
	if IsFile(fileName) {
		return fileName, errors.New("文件已经存在")
	}
	// 创建csv文件
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return fileName, err
	}
	defer f.Close()

	csvFile := csv.NewWriter(f)
	defer csvFile.Flush()

	// 处理字段
	err = csvFile.Write(title)
	if err != nil {
		return fileName, err
	}
	// 处理数据
	for _, v := range data {
		var subValueList []string
		for _, value := range v {
			subValueList = append(subValueList, fmt.Sprintf("%v", value))
		}
		valueList = append(valueList, subValueList)
	}

	err = csvFile.WriteAll(valueList)
	if err != nil {
		return fileName, err
	}
	return fileName, nil
}

func ExportJson(data []byte, filename string) (string, error) {
	if IsFile(filename) {
		return filename, errors.New("文件已经存在")
	}
	// 创建json文件
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return filename, err
	}
	defer f.Close()

	if err != nil {
		return filename, err
	}

	_, err = f.Write(data)

	if err != nil {
		return filename, err
	}
	return filename, nil
}
