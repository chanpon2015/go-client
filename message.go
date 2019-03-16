package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "test":
		payload = "ok"
		sendNotification()
	case "download":
		file, err := os.Open("/test/sample.tsv")
		if err != nil {
			payload = err
		}
		b, _ := ioutil.ReadAll(file)
		payload = base64.StdEncoding.EncodeToString(b)
	default:
	}
	return
}
