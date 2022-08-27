# pmccron
定时任务调度器

#### 使用方式

```
go get github.com/PandaManPMC/pmccron@v1.0.2
```

```go
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

	scheduler.Running()
	time.Sleep(24 * time.Hour)
}
```


