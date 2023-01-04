package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 165. 比较版本号
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	for i:=0; i < len(v1) || i < len(v2); i ++ {
		x,y := 0,0
		if i < len(v1) {
			x, _= strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}

// 151. 反转字符串中的单词
func reverseWords(s string) string {
	lenth := len(s)
	if lenth <= 1 {
		return s
	}
	res := ""
	for i:=0; i < lenth ; {
		word := ""
		j := i
		// 读取单词
		for j < lenth && s[j] != ' '{
			j ++
		}
		word = s[i:j]
		// 跳过空格
		for j < lenth && s[j] == ' ' {
			j ++
		}
		// 拼接单词  注意这里
		if word != "" {
			res = " " + word + res
		}

		i = j
	}
	return res[1:]
}

func main(){

	fmt.Print()
}