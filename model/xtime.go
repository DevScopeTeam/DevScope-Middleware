package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 定义Time，继承自time.Time
type XTime struct {
	time.Time
}

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *XTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = XTime{now}
	return
}

func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(timeFormart))
	return []byte(output), nil
}

// Before
func (t XTime) Before(t2 time.Time) bool {
	return time.Time(t.Time).Before(time.Time(t2))
}

// After
func (t XTime) After(t2 time.Time) bool {
	return time.Time(t.Time).After(time.Time(t2))
}

// Parse
func (t XTime) Parse(value string) (XTime, error) {
	// 统一为东八区
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, err := time.ParseInLocation(timeFormart, value, loc)
	return XTime{tt}, err
}

// 适配 gorm

// 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
