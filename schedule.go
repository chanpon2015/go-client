package main

import (
	"math"

	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/robfig/cron"
	"github.com/shirou/gopsutil/cpu"
)

func cronStart() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		per, _ := cpu.Percent(0, false)
		bootstrap.SendMessage(w, "using_cpu", round(per[0]))
	})
	go c.Start()
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
