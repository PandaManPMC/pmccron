package pmccron

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"sync"
	"time"
)

//  author: laoniqiu
//  since: 2022/8/27
//  desc: pmccron 调度器

var running bool

func runScheduler() {
	go func() {
		defer func() {
			err := recover()
			if nil != schedulerInstance.logError {
				schedulerInstance.logError("", err)
			}
			runScheduler()
		}()

		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			if running {
				continue
			}
			running = true
			startTask()
			running = false
		}
	}()
}

//	startTask 遍历定时任务，检测是否需要启动
func startTask() {
	schedulerInstance.lock.Lock()
	defer schedulerInstance.lock.Unlock()
	for i := 0; i < len(schedulerInstance.taskList); i++ {
		task := schedulerInstance.taskList[i]
		if task.deleted {
			continue
		}
		now := time.Now()
		//schedulerInstance.logInfo(fmt.Sprintf("检查定时任务%d，cron=%s，下次执行时间是%s，现在时间是%s", task.sn, task.cron, time.Unix(task.nextTime, 0).String(), now.String()))
		if task.nextTime < now.Unix() {
			// 表示到了执行时间，执行定时任务
			go task.fun()
			// 计算下次执行时间
			cpr := cronexpr.MustParse(task.cron)
			nextTime := cpr.Next(time.Now())
			task.nextTime = nextTime.Unix()
			schedulerInstance.logInfo(fmt.Sprintf("定时任务%d执行,cron=%s,下次执行时间计算%s", task.sn, task.cron, nextTime.String()))
		}
	}
}

type scheduler struct {
	lock     sync.Mutex
	running  bool
	nextSn   uint
	taskList []*schedulerTask
	logInfo  func(msg string)
	logError func(msg string, err interface{})
}

var schedulerInstance *scheduler = nil

//	InitSchedulerSingle 定时任务调度器实例化【单例】
//	logInfo  func(msg string)	info 日志输出函数
//	logError func(msg string, err interface{}) err 日志输出函数
func InitSchedulerSingle(logInfo func(msg string), logError func(msg string, err interface{})) *scheduler {
	if nil == logInfo {
		return nil
	}
	if nil == logError {
		return nil
	}
	if nil == schedulerInstance {
		tl := make([]*schedulerTask, 0)
		schedulerInstance = &scheduler{
			lock:     sync.Mutex{},
			running:  false,
			nextSn:   0,
			taskList: tl,
			logInfo:  logInfo,
			logError: logError,
		}
	}
	return schedulerInstance
}

//	Running 启动定时任务调度器
func (instance *scheduler) Running() bool {
	instance.lock.Lock()
	defer instance.lock.Unlock()
	if instance.running {
		return true
	}
	instance.running = true
	runScheduler()
	return instance.running
}

//	add 添加定时任务
//	cron string 完整的 cron 表达式
//	task func()	任务
//	uint 任务编号，0 则表示失败
func (instance *scheduler) add(cron string, task func()) uint {
	instance.lock.Lock()
	defer instance.lock.Unlock()
	cpr, err := cronexpr.Parse(cron)
	if nil != err {
		return 0
	}
	next := cpr.Next(time.Now())
	instance.nextSn++
	st := schedulerTask{
		sn:       instance.nextSn,
		deleted:  false,
		cron:     cron,
		fun:      task,
		nextTime: next.Unix(),
	}
	instance.taskList = append(instance.taskList, &st)
	instance.logInfo(fmt.Sprintf("定时任务%d，cron=%s，下次执行时间%s", st.sn, st.cron, next.String()))
	return st.sn
}

//	DayHour 增加每天执行一次的任务
//	hour string	每天执行的时间，如 08 则在每日 8 时执行一次
//	task func()	任务
func (instance *scheduler) DayHour(hour string, task func()) uint {
	if nil == task {
		return 0
	}
	return instance.add(fmt.Sprintf("0 %s * * *", hour), task)
}

//	Minute 每到这个分执行一次
//	minute string 每个小时执行的分时，如 5 则在每时的第 5 分执行一次
//	task func() 任务
func (instance *scheduler) Minute(minute string, task func()) uint {
	if nil == task {
		return 0
	}
	return instance.add(fmt.Sprintf("%s * * * *", minute), task)
}
