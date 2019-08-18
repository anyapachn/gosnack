package main

import (
	"fmt"
	"math"
)

func mod(x, y float64) float64 {
	float := math.Mod(x, y)
	return float
}

func div(hi, y int) int {
	a := hi / y
	return a
}

func roman_num(num1, num2 int) (cho1, cho2 string) {
	var ch1, ch2 string
	switch num1 {
	case 1:
		ch1 = "X"
	case 2:
		ch1 = "XX"
	case 3:
		ch1 = "XXX"
	case 4:
		ch1 = "XL"
	case 5:
		ch1 = "L"
	case 6:
		ch1 = "LX"
	case 7:
		ch1 = "LXX"
	case 8:
		ch1 = "LXXX"
	case 9:
		ch1 = "XC"
	case 10:
		ch1 = "C"
	}
	switch num2 {
	case 1:
		ch2 = "I"
	case 2:
		ch2 = "II"
	case 3:
		ch2 = "III"
	case 4:
		ch2 = "IV"
	case 5:
		ch2 = "V"
	case 6:
		ch2 = "VI"
	case 7:
		ch2 = "VII"
	case 8:
		ch2 = "VIII"
	case 9:
		ch2 = "IX"
	case 0:
		ch2 = ""
	}
	return ch1, ch2
}

func main() {
	fmt.Println("Start print Roman numeral")
	var rnum string
	for i := 1; i <= 100; i++ {
		num1 := float64(i)
		x := mod(num1, 10)
		no1 := div(i, 10)
		//fmt.Println(no1, x)
		no2 := int(x)
		rnum1, rnum2 := roman_num(no1, no2)
		rnum = rnum1 + rnum2
		fmt.Println(i, " = ", rnum)
	}
}
