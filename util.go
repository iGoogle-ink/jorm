package jorm

import (
	"strconv"
	"strings"
	"time"
)

const (
	null       string = ""
	TimeLayout string = "2006-01-02 15:04:05"
	DateLayout string = "2006-01-02"
)

//解析时间
func ParseDateTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

//格式化Datetime
func FormatDateTime(timeStr string) (formatTime string) {
	//2019-01-04T15:40:00Z
	//2019-01-18 20:51:30+08:00
	if timeStr == "" {
		return ""
	}
	replace := strings.Replace(timeStr, "T", " ", 1)
	formatTime = replace[:19]
	return
}

//格式化
func FormatDate(dateStr string) (formatDate string) {
	//2020-12-30T00:00:00+08:00
	if dateStr == "" {
		return ""
	}
	split := strings.Split(dateStr, "T")
	formatDate = split[0]
	return
}

//字符串转Float
func String2Float(floatStr string) (floatNum float64) {
	floatNum, _ = strconv.ParseFloat(floatStr, 64)
	return
}

//Float64转字符串
//    floatNum：float64数字
//    prec：精度位数（不传则默认float数字精度）
func Float64ToString(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
	return
}

//Float32转字符串
//    floatNum：float32数字
//    prec：精度位数（不传则默认float数字精度）
func Float32ToString(floatNum float32, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(float64(floatNum), 'f', prec[0], 32)
		return
	}
	floatStr = strconv.FormatFloat(float64(floatNum), 'f', -1, 32)
	return
}

//字符串转Int
func String2Int(intStr string) (intNum int) {
	intNum, _ = strconv.Atoi(intStr)
	return
}

//字符串转Int64
func String2Int64(intStr string) (int64Num int64) {
	intNum, _ := strconv.Atoi(intStr)
	int64Num = int64(intNum)
	return
}

//Int转字符串
func Int2String(intNum int) (intStr string) {
	intStr = strconv.Itoa(intNum)
	return
}

//Int64转字符串
func Int642String(intNum int64) (int64Str string) {
	int64Str = strconv.FormatInt(intNum, 10)
	return
}
