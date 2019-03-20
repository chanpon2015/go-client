package excel

import (
	"encoding/csv"
	"os"
	"path/filepath"

	"github.com/tealeg/xlsx"
)

func ExportTsvFile(xlsxPath, exportDir string) ([]string, error) {
	file, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		return nil, err
	}
	var filePaths []string
	for _, sheet := range file.Sheets {
		var sheetData [][]string
		for _, row := range sheet.Rows {
			var rowData []string
			for _, cell := range row.Cells {
				rowData = append(rowData, cell.String())
			}
			sheetData = append(sheetData, rowData)
		}
		filePath, err := writerTsv(exportDir, sheet.Name, sheetData)
		if err != nil {
			return nil, err
		}
		filePaths = append(filePaths, filePath)
	}
	return filePaths, nil
}

func writerTsv(exportDir, name string, data [][]string) (string, error) {
	file, err := os.OpenFile(
		filepath.Join(
			exportDir,
			name+".tsv",
		), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return "", err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	for i, d := range data {
		writer.Write(d)
		if i%100 == 0 {
			writer.Flush()
		}
	}
	writer.Flush()
	return file.Name(), nil
}
