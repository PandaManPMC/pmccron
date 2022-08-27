package pmccron

//  author: laoniqiu
//  since: 2022/8/27
//  desc: pmccron

type schedulerTask struct {
	sn       uint   // 编号
	deleted  bool   // 删除
	cron     string // 表达式
	fun      func() // 任务函数
	nextTime int64  // 下次执行时间戳
}
