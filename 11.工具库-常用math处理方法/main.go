package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

// 36进制与10进制互转
func convert36to10()  {
	data := strconv.FormatUint(212121212, 36) //10进制转36进制
	fmt.Println(data)
	fmt.Println(strconv.ParseUint(data, 36, 64)) //36进制转10进制
}


// 任意进制转10进制
func anyToDecimalMap(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkeyMap(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}

// map根据value找key
func findkeyMap(in string) int {
	result := -1
	for k, v := range tenToAnyMap {
		if in == v {
			result = k
		}
	}
	return result
}

var tenToAnyMap map[int]string = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z"}

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
