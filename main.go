package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
)

var (
	w *astilectron.Window
)

func main() {
	astilog.Debug("start")
	err := bootstrap.Run(
		bootstrap.Options{
			Asset:    Asset,
			AssetDir: AssetDir,
			AstilectronOptions: astilectron.Options{
				AppName:            "go-vue-typescript-electron",
				AppIconDefaultPath: "resources/icon.png",
			},
			Debug: true,
			MenuOptions: []*astilectron.MenuItemOptions{
				{
					Label: astilectron.PtrStr("File"),
					SubMenu: []*astilectron.MenuItemOptions{
						{
							Label: astilectron.PtrStr("About"),
							OnClick: func(e astilectron.Event) (deleteListener bool) {
								return
							},
						},
					},
				},
			},
			OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
				w = ws[0]
				// cronStart()
				// createNotification()
				return nil
			},
			RestoreAssets: RestoreAssets,
			Windows: []*bootstrap.Window{
				{
					Homepage:       "index.html",
					MessageHandler: handleMessages,
					Options: &astilectron.WindowOptions{
						Center: astilectron.PtrBool(true),
						Height: astilectron.PtrInt(500),
						Width:  astilectron.PtrInt(500),
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
}
