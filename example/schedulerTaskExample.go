package main

import (
	"fmt"
	"pmccron/pmccron"
	"sync"
	"time"
)

//  author: laoniqiu
//  since: 2022/8/27
//  desc: example

func main() {

	wg := sync.WaitGroup{}
	wg.Add(2)

	scheduler := pmccron.InitSchedulerSingle(func(msg string) {
		fmt.Println(msg)
	}, func(msg string, err interface{}) {
		fmt.Printf("%s----------%v", msg, err)
	})
	scheduler.DayHour("08", func() {
		fmt.Println("执行每天 08 点定时任务")
		wg.Done()
	})
	scheduler.DayHour("16", func() {
		fmt.Println("执行每天 16 点定时任务")
		wg.Done()
	})
	scheduler.Minute("15", func() {
		fmt.Println("执行每 15 分钟执行一次的定时任务")
	})

	scheduler.Running()
	//wg.Wait()
	time.Sleep(24 * time.Hour)
}
