package main

import (
	"github.com/asticode/go-astilectron"
)

var (
	a *astilectron.Astilectron
)

func createNotification() {
	a, _ = astilectron.New(astilectron.Options{
		AppName: "test",
	})
	defer a.Close()
	a.Start()
	a.Wait()
}

func sendNotification() {
	n := a.NewNotification(&astilectron.NotificationOptions{
		Body:  "test",
		Icon:  "/resources/icon.png",
		Title: "test",
	})
	n.On(astilectron.EventNameNotificationEventClicked, func(e astilectron.Event) (deleteListener bool) {
		return
	})
	n.Create()
	n.Show()
}
