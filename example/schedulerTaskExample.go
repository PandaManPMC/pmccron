package main

import (
	"fmt"
	"github.com/PandaManPMC/pmccron"
	"time"
)

//  author: laoniqiu
//  since: 2022/8/27
//  desc: example

func main() {
	scheduler := pmccron.InitSchedulerSingle(func(msg string) {
		fmt.Println(msg)
	}, func(msg string, err interface{}) {
		fmt.Printf("%s----------%v", msg, err)
	})
	scheduler.DayHour("08", func() {
		fmt.Println("执行每天 08 点定时任务")
	})
	scheduler.DayHour("16", func() {
		fmt.Println("执行每天 16 点定时任务")
	})
	scheduler.Minute("15", func() {
		fmt.Println("执行每时的第 15 分的定时任务")
	})

	scheduler.Cron("0 8 18 28 * ? *", func() {
		fmt.Println("根据表达式执行")
	})

	count := 0
	scheduler.Cron("0 * * * * ? *", func() {
		fmt.Println("每 1 分执行一次的定时任务")
		count += 1
		if 1 == count {
			panic("异常测试")
		}
	})

	scheduler.CronByName("每小时的1分钟0秒执行", pmccron.MinuteOneInEveryHour, func() {
		fmt.Println("根据表达式执行")
	})

	scheduler.Running()
	time.Sleep(24 * time.Hour)

}
