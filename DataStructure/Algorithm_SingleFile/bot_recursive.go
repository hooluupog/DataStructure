/*
//题目
//有一个机器人，给出一串指令，L左转 R右转，F前进一步，B后退一步，求最后机器人的坐标。
//最开始，机器人位于 0 0，方向为正Y。可以输入重复指令n ：如 R2(LF) 等于指令 RLFLF。
// 示例输入：
//   FRB
//   R2(B2(LF)BF2(BF))FBF
//   FRRB
//   FRRBLFRB
//   FRRRBLLFRBLF
//   R2(B2(LF))F
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var x, y int
	var input string
	fmt.Scan(&input)
	cmds := getCmd(input)
	for _, cmd := range cmds {
		switch string(cmd) {
		case "F":
			y++
		case "B":
			y--
		case "L":
			x--
		case "R":
			x++
		}
	}
	fmt.Println("(", x, ",", y, ")")
}

var pos int = 0

// recursive.
func getCmd(s string) string {
	if pos < len(s) {
		a := string(s[pos])
		switch {
		case a == ")":
			return ""
		case a == "(":
			pos++
			return getCmd(s)
		case a >= "0" && a <= "9":
			repeat := ""
			for a >= "0" && a <= "9" {
				repeat += a
				pos++
				a = string(s[pos])
			}
			count, _ := strconv.Atoi(repeat)
			subCmd := strings.Repeat(getCmd(s), count)
			pos++
			return subCmd + getCmd(s)
		default:
			pos++
			return a + getCmd(s)
		}
	}
	return ""
}
