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

//Non-recursive.
func getCmd(s string) (res []byte) {
	cmd := make([]byte, 0)
	for _, v := range s {
		a := string(v)
		if a == ")" { //  pop out of stack cmd
			repeatCmd := ""    // repeat times.
			subCmd := ""       // repeated commands.
			repeat := false    // start recording repeat number.
			for len(cmd) > 0 { // translate repeated commands.
				b := string(cmd[len(cmd)-1])
				if b == "(" {
					cmd = cmd[:len(cmd)-1]
					repeat = true
					continue
				}
				for b >= "0" && b <= "9" {
					repeatCmd += b
					cmd = cmd[:len(cmd)-1]
					if len(cmd) == 0 {
						break
					}
					b = string(cmd[len(cmd)-1])
				}
				if repeat {
					count, _ := strconv.Atoi(repeatCmd)
					subCmd = strings.Repeat(subCmd, count)
					for i := len(subCmd) - 1; i >= 0; i-- {
						cmd = append(cmd, subCmd[i])
					}
					repeat = false
					break
				}
				subCmd += string(cmd[len(cmd)-1])
				cmd = cmd[:len(cmd)-1]
			}
			continue
		}
		// push into stack cmd
		cmd = append(cmd, byte(v))
	}
	return cmd
}
