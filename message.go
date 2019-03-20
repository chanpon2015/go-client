package main

import (
	"encoding/json"

	"github.com/chanpon2015/go-client/excel"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

type ExcelToTsvPayload struct {
	XlsxPath  string `json:"xlsx_path"`
	ExportDir string `json:"export_dir"`
}

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	switch m.Name {
	case "excel_to_tsv":
		var p ExcelToTsvPayload
		if err := json.Unmarshal(m.Payload, &p); err != nil {
			return nil, err
		}
		excel.ExportTsvFile(p.XlsxPath, p.ExportDir)
		return &p, nil
	default:
		return nil, nil
	}
}
