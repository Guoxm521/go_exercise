package model

import (
	"example.com/m/v2/Project/model/db"
	"example.com/m/v2/Project/spider"
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
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
		_searchList := make([]*spider.SearchStruct, 0)
		//语言列表
		_language, _err := db.NewGithubLanguage().SearchAll()
		//时间段
		_timeSince, _err := db.NewGithubSince().SearchAll()
		for _, _v1 := range _language {
			for _, _v2 := range _timeSince {
				_item := new(spider.SearchStruct)
				_item.Since = _v2.Type
				_item.SinceType = _v2.SinceType
				_item.Language = _v1.Name
				_item.LanguageType = _v1.Type
				_searchList = append(_searchList, _item)
			}
		}
		fmt.Println("========err========", _err)
		fmt.Println("-==================准备只选")
		//for _, _v1 := range _searchList {
		//	time.Sleep(10 * time.Second)
		//	fmt.Println("=========================", _v1)
		//	for {
		//		github := new(spider.GithubTrending)
		//		data, _err := github.NewCollector(_v1).SpiderGithub()
		//		if _err != "" {
		//			fmt.Println("===========我是爬虫中的错误===============，", _err)
		//			time.Sleep(10 * time.Second)
		//		} else {
		//			fmt.Println("=================结果==================", data)
		//			break
		//		}
		//	}
		//}
		c := cron.New(cron.WithSeconds())
		//c.AddFunc("@every 10s", func() {
		//})
		c.AddFunc("00 50 22 * * ?", func() {
			fmt.Println("=== 每天23点50汇总当日进出客流量推送 ===")
			for _, _v1 := range _searchList {
				time.Sleep(120 * time.Second)
				for {
					github := new(spider.GithubTrending)
					data, _err := github.NewCollector(_v1).SpiderGithub()
					if _err != "" {
						fmt.Println("===========我是爬虫中的错误===============，", _err)
						time.Sleep(60 * time.Second)
					} else {
						fmt.Println("=================结果==================", data)
						break
					}
				}
			}
		})
		//c.AddFunc("@hourly", func() { fmt.Println("每小时执行") })
		//c.AddFunc("@every 1h30m", func() { fmt.Println("每小时30分执行") })
		//c.AddFunc("@daily", func() {fmt.Println("===每天午夜执行===")})
		//c.AddFunc("@weekly", func() { fmt.Println("每周运行一次，周六/周日之间的午夜") })
		//c.AddFunc("@monthly", func() { fmt.Println("每月运行一次，午夜，第一个月") })
		c.Start()
		select {}
	}
}
