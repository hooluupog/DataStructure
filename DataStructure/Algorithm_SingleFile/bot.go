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
	cmds, _ := getCmd(input, 0)
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
func getCmd(s string, pos int) (res string, po int) {
	repeat := ""
	subCmd := ""
	i := 0
	for i = pos; i < len(s); i++ {
		a := string(s[i])
		if a == "(" {
			subCmd, i = getCmd(s, i+1)
			count, _ := strconv.Atoi(repeat)
			subCmd = strings.Repeat(subCmd, count)
			res += subCmd
			continue
		}
		if a == ")" {
			po = i
			return res, po
		}
		if a >= "0" && a <= "9" {
			repeat += a
			continue
		}
		res += a
	}
	po = i
	return res, po
}
