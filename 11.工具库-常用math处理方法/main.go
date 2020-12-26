package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * @Author: SimonWang00
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2020/12/26 15:23
 */


//f 保留2 位小数 %.2f，默认四舍五入
//Decimal(22.66666666666,"%.0f") 四舍五入是23
//Decimal(22.222222222,"%.2f") 两位小数是22.22
func Decimal(value float64,f string) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf(f, value), 64)
	return value
}

//waitSecond 该用户等待了多少秒
// input 1 output 4
func getRatingStep(waitSecond int64) int {
	var (
		step             = 1.3
		baseStep float64 = 3
		maxStep  float64 = 100
	)
	u := math.Pow(float64(waitSecond), step)		// 1, x的y次方
	u = u + baseStep								// 4
	u = math.Round(u)								// 4
	u = math.Min(u, maxStep) 						// 等待时间越长，rating 区间越大
	return int(u)
}

func main() {
	step := getRatingStep(2)
	print(step)
}
