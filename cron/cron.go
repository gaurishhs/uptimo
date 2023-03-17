package cron

import "github.com/robfig/cron/v3"

func StartCronJobs() {
	c := cron.New()
	c.AddFunc("0/30 0/1 * 1/1 * ? *", func() {
		// Do something here
	})
	c.Start()
}
