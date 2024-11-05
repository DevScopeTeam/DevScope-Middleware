package utils

import (
	"DevScope-Middleware/model"
	"time"
)

// utils.GetTimeNow 获取当前时间
func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetXTimeNow() model.XTime {
	return model.XTime{Time: time.Now()}
}

// GetDateNow 获取当前日期
func GetDateNow() string {
	return time.Now().Format("2006-01-02")
}

// utils.GetTimestamp 获取当前时间戳
func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}
