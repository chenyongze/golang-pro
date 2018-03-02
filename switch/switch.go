//switch 任务
//author:yongze.chen
package main

import (
	"fmt"
)

func main() {
	fmt.Println("switch start")
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Printf("你的等级是 %s\n", grade)

}
