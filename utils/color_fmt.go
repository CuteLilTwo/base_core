package utils

import (
	"fmt"
)

//带颜色输出fmt方法 backColor：背景颜色   fontColor：字体颜色   fontStyle：字体样式   content：内容
func ColorFmt(backColor string, fontColor string, fontStyle string, Enter bool, content ...interface{}) {
	var backColorInt int
	var fontColorInt int
	var fontStyleInt int

	switch fontStyle {
	case "默认":
		backColorInt = 0
	case "高亮":
		backColorInt = 1
	case "下划线":
		backColorInt = 4
	case "闪烁":
		backColorInt = 5
	case "反白":
		backColorInt = 7
	case "不可见":
		backColorInt = 8
	}

	switch backColor {
	case "黑色":
		backColorInt = 40
	case "红色":
		backColorInt = 41
	case "绿色":
		backColorInt = 42
	case "黄色":
		backColorInt = 43
	case "蓝色":
		backColorInt = 44
	case "紫红色":
		backColorInt = 45
	case "青蓝色":
		backColorInt = 46
	case "白色":
		backColorInt = 47
	}

	switch fontColor {
	case "黑色":
		fontColorInt = 30
	case "红色":
		fontColorInt = 31
	case "绿色":
		fontColorInt = 32
	case "黄色":
		fontColorInt = 33
	case "蓝色":
		fontColorInt = 34
	case "紫红色":
		fontColorInt = 35
	case "青蓝色":
		fontColorInt = 36
	case "白色":
		fontColorInt = 37
	}
	if Enter == true {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m\n", 0x1B, fontStyleInt, backColorInt, fontColorInt, content, 0x1B)
	} else {
		fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1B, fontStyleInt, backColorInt, fontColorInt, content, 0x1B)
	}
}
