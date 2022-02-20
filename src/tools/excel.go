package tools

import (
	excelize "github.com/xuri/excelize/v2"
)

func ReadExcel(path string, sheet string) (item [][]string, err error) {
	f, err := excelize.OpenFile(path)
	if err != nil {

		return nil, err
	}
	item, _ = f.GetRows(sheet)
	return item, nil
}
