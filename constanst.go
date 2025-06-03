package pmccron

const (
	// MinuteOneInEveryHour 每小时的第1分0秒执行
	MinuteOneInEveryHour = "0 1 * * * * *"

	// FiveMinuteForZero 从 0 开始每 5 分钟 0 秒执行
	FiveMinuteForZero = "0 0/5 * * * * *"
)
