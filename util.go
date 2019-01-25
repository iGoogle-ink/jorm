//==================================
//  * Name：Jerry
//  * Tel：18017448610
//  * DateTime：2019/1/23 14:30
//==================================
package jorm

import "strconv"

const (
	null string = ""
)

//字符串转Float
func String2Float(floatStr string) (floatNum float64) {
	floatNum, _ = strconv.ParseFloat(floatStr, 64)
	return
}

//Float转字符串
//    floatNum：float数字
//    prec：精度位数（不传则默认float数字精度）
func Float2String(floatNum float64, prec ...int) (floatStr string) {
	if len(prec) > 0 {
		floatStr = strconv.FormatFloat(floatNum, 'f', prec[0], 64)
		return
	}
	floatStr = strconv.FormatFloat(floatNum, 'f', -1, 64)
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
