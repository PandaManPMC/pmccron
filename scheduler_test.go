package pmccron

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"testing"
	"time"
)

//  author: laoniqiu
//  since: 2022/8/27
//  desc: pmccron

func TestFunc(t *testing.T) {
	scheduler := InitSchedulerSingle(func(msg string) {
		fmt.Println(msg)
	}, func(msg string, err interface{}) {
		fmt.Println(msg)
		fmt.Println(err)
	})
	t.Log(nil == scheduler.logInfo)
	t.Log(nil == scheduler.logError)
}

func TestCron(t *testing.T) {
	c := "0 08 * * *"
	m := cronexpr.MustParse(c)
	t.Log(m.Next(time.Now()))

	t.Log(cronexpr.MustParse("40 14 * * *").Next(time.Now()))

	//	每月 1 号 上午 8 点 执行
	t.Log(cronexpr.MustParse("0 0 8 1 * ? *").Next(time.Now()))

}
