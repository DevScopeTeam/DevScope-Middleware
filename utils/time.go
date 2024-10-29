package utils

import (
	"time"
)

// GetTimeNow 获取当前时间
func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetDateNow 获取当前日期
func GetDateNow() string {
	return time.Now().Format("2006-01-02")
}

// GetTimestamp 获取当前时间戳
func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}
