package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 { // O(lgn) complexity.
	if n == 0 {
		return 1
	}
	mul, e := float64(1), 0
	if n < 0 {
		if n == -(1 << 31) { // 最小负数转为正数溢出的情况处理
			e = (1 << 31) - 1
		} else {
			e = -n
		}
	} else {
		e = n
	}
	for i, val := 0, x; i < 31; i++ { // n is 32bits and the highest bit is sign.
		if e&0x1 == 1 { // 当前位为1，则相乘
			mul = mul * val
		}
		val = val * val
		e = e >> 1
	}
	if n > 0 {
		return mul
	} else {
		if n == -(1 << 31) {
			mul = mul * x
		}
		return float64(1) / mul
	}
}

func main() {
	fmt.Println(myPow(2.00000, -2147483648))
	fmt.Println(myPow(-2.00000, -5))
}
