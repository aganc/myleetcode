package main

import (
	"fmt"
	"math"
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

// 子集 用回溯法，注意golanappend的时候用copy，不然会右错乱
func subsets(nums []int) [][]int {
	var res [][]int
	vis := map[int]bool{}
	var trackBack func([]int, int)
	trackBack = func(path []int, j int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		for i := j; i < len(nums); i ++ {
			if vis[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			vis[nums[i]] = true
			trackBack(path, i)
			path = path[:len(path)-1]
			vis[nums[i]] = false
		}
	}
	trackBack([]int{}, 0)
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	res := 0
	var traceBack func (node *TreeNode, result int)
	traceBack = func (node * TreeNode, result int) {
		if node == nil {
			return
		}
		result = result * 10 + node.Val
		if node.Left == nil && node.Right == nil {
			res += result
			// fmt.Printf("result = %d, res = %d\n", result, res)
			return
		}
		// fmt.Printf("result = %d\n", result)
		if node.Left != nil {
			traceBack(node.Left, result)
		}
		if node.Right != nil {
			traceBack(node.Right, result)
		}
	}
	traceBack(root, 0)
	return res
}

// 递归实现
func isValidBST(root *TreeNode) bool {
	var digui func(node *TreeNode, left int, Right int) bool
	digui = func(node *TreeNode, Left int, Right int) bool {
		if node == nil {
			return true
		}
		if node.Val >= Right || node.Val <= Left {
			return false
		}
		if digui(node.Left, Left, node.Val) == false || digui(node.Right, node.Val, Right) == false {
			return false
		}
		return true
	}
	return digui(root, math.MinInt64, math.MaxInt64)
}

func main(){

	// root := []int{1, 2, 3}


	fmt.Print()
}