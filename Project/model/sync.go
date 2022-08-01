package model

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

type Sync struct {
}

func NewSync() *Sync {
	return &Sync{}
}

//使用一个协程
func GoFunc(f func()) {
	go func() {
		f()
	}()
}

func (that *Sync) Cron() func() {
	fmt.Println("========================================")
	return func() {
		fmt.Println("=====================定时任务===================")
		c := cron.New(cron.WithSeconds())
		c.AddFunc("@every 20s", func() {
			fmt.Println("123123123123213333")
		})
		//c.AddFunc("00 50 23 * * ?", func() {
		//	fmt.Println("=== 每天23点50汇总当日进出客流量推送 ===")
		//
		//})
		//c.AddFunc("@hourly", func() { fmt.Println("每小时执行") })
		//c.AddFunc("@every 1h30m", func() { fmt.Println("每小时30分执行") })
		//c.AddFunc("@daily", func() {fmt.Println("===每天午夜执行===")})
		//c.AddFunc("@weekly", func() { fmt.Println("每周运行一次，周六/周日之间的午夜") })
		//c.AddFunc("@monthly", func() { fmt.Println("每月运行一次，午夜，第一个月") })
		c.Start()
		select {}
	}
}
