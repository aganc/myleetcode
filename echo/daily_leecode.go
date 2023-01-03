package main

import (
	"fmt"
	"sort"
)

// 岛屿数量，使用dfs遍历
func numIslands(grid [][]byte) int {

	line := len(grid)
	col := len(grid[0])
	if line == 0 && col == 0 {
		return 0
	}
	var dfs func(int, int, int, int, [][]byte)
	dfs = func(x int, y int, line int, col int, grid [][]byte){
		if x<0 || x >= line || y <0 || y >= col || grid[x][y] != '1'{
			return
		}
		grid[x][y] = '2'
		dfs(x + 1, y, line, col, grid)
		dfs(x - 1, y, line, col, grid)
		dfs(x, y - 1, line, col, grid)
		dfs(x, y + 1, line, col, grid)
	}

	result := 0
	for i := 0; i < line; i ++ {
		for j := 0; j < col; j ++{
			if grid[i][j] == '1'{
				result ++
				dfs(i, j, line, col, grid)
			}
		}
	}
	return result
}


// 最长回文子串，使用动态规划
// dp[i][j]表示 i ...j 是否回文。
// dp[i][j]  =   true    i==j       false  dp[i] != dp[j]
// dp[i]==dp[j]   j-i < 3 true  esle  dp[i][j] = dp[i+1][j-1]
func longestPalindrome(s string) string {
	if s == "" || len(s) == 1{
		return s
	}
	lenth := len(s)
	dp := make([][]bool, lenth)
	for i := range dp{
		dp[i] = make([]bool, lenth)
		dp[i][i] = true
	}

	left, maxLen := 0, 1
	// 先遍历子串的长度
	for i:=2; i<lenth+1; i++{
		for j:=0; j<lenth;j++{
			k := j+i-1
			if k >= lenth{
				break
			}
			if s[j] != s[k]{
				dp[j][k] = false
			} else if k - j < 3{
				dp[j][k] = true
			} else {
				dp[j][k] = dp[j+1][k-1]
			}
			if dp[j][k] && i > maxLen{
				left = j
				maxLen = i
			}
		}
	}
	return s[left:left+maxLen]
}


// 层序遍历的变种，添加倒序flag， 锯齿遍历二叉树
/**
* Definition for a binary tree node.
 */
type TreeNode struct {
	Val int
   	Left *TreeNode
   	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil{
		return result
	}
	queue := []*TreeNode{root}
	flag := false
	for {
		tmpval := []int{}
		q := queue
		queue = nil
		for _,node := range q{
			tmpval = append(tmpval, node.Val)
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
		}
		if flag == true{
			for i, n := 0, len(tmpval); i < n/2; i++ {
				tmpval[i], tmpval[n-1-i] = tmpval[n-1-i], tmpval[i]
			}
		}
		result = append(result, tmpval)
		flag = !flag
		if len(queue) == 0{
			break
		}
	}

	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val{
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil{
		return right
	} else if right == nil{
		return left
	}
	return root
}

// 全排列，用回溯法，注意还原，回溯是一个专题
func permute(nums []int) [][]int {
	var res [][]int
	visited := map[int]bool{}

	var backTrack func(path []int)
	backTrack = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for _, n := range nums{
			if visited[n]{
				continue
			}
			path = append(path, n)
			visited[n] = true
			backTrack(path)
			path = path[: len(path) - 1]
			visited[n] = false
		}
	}
	backTrack([]int{})
	return res
}
// 螺旋矩阵，采用模拟法，注意边界值x --  ，y ++的处理
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := []int{}
	if m == 0 && n ==0 {
		return res
	}
	x, y := 0, -1
	right, down, left, up := n-1, m-1, 0, 0
	for left <= right && up <= down{
		// 向右
		for y++; y <= right; y ++{
			res = append(res, matrix[x][y])
		}
		up ++
		y --
		//fmt.Printf("%v,%v,%v,%v, res = %v\n",right,down,left,up,res)
		// 向下
		for x++; x <= down; x ++{
			res = append(res, matrix[x][y])
		}
		right --
		x --
		if left <= right && up <= down{
			// 向左
			for y--;y >= left; y --{
				res= append(res, matrix[x][y])
			}
			down --
			y ++
			//fmt.Printf("%v,%v,%v,%v, res = %v\n",right,down,left,up,res)
			// 向上
			for x--; x >= up; x --{
				res = append(res,matrix[x][y])
			}
			left ++
			x ++
		}
		// fmt.Printf("%v,%v,%v,%v, res = %v\n",right,down,left,up,res)
		//fmt.Printf("%v,%v,%v,%v, res = %v\n",right,down,left,up,res)
	}
	return res
}


// 反转链表2  穿针引线法，一次遍历，需要用到三个变量
/** */
 // Definition for singly-linked list.
 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right{
		return head
	}
	res := &ListNode{Val:-1}
	res.Next = head
	pre := res
	for i:=0; i< left-1; i ++{
		pre = pre.Next
	}
	cur := pre.Next
	for i:=0; i < right-left; i++{
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

	}
	return res.Next
}

 // 环形链表，最简单的办法是哈希表法
func detectCycle(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	setMap := make(map[*ListNode]struct{})
	for ; head != nil; head = head.Next{
		// fmt.Printf("map = %v", setMap)
		if _,ok := setMap[head]; ok{
			return head
		}
		setMap[head] = struct{}{}
	}
	return head
}

// 300. 最长递增子序列  动态规划，要学习下这种写法
func lengthOfLIS(nums []int) (ans int) {
	dp := make([]int, len(nums))
	for i := range dp {dp[i] = 1} // 初始化，dp[i]表示递增子序列的长度
	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] { // nums[i] 可加入递增子序列
				dp[i] = max(dp[i], dp[j] + 1) // 有很多递增子序列，求这些最长递增子序列的最大长度
			}
		}
	}
	for _, d := range dp {ans = max(ans, d)}
	return
}

func max(a, b int) int {if a > b {return a}; return b}


 // 20 有效得括号，模拟栈，先进后出
func isValid(s string) bool {
	sLen := len(s)
	if sLen % 2 != 0{
		return false
	}
	setMap := map[byte]byte{
		']':'[',
		'}':'{',
		')':'(',
	}
	var tmp []byte
	for i := 0; i < sLen; i ++{
		if s[i] == '[' || s[i] == '(' || s[i] == '{'{
			tmp = append(tmp, s[i])
			continue
		}
		if len(tmp) <= 0 || tmp[len(tmp) - 1] != setMap[s[i]]{
			return false
		}
		tmp = tmp[:len(tmp) - 1]
	}
	return len(tmp) == 0
}


 // 重排链表，最基本得切片存储对应所有值
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	if head == nil{
		return
	}
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next{
		nodes = append(nodes, node)
	}
	// fmt.Printf("nodes: %v", nodes)
	i ,j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}

		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}

// 199 二叉树的右视图，属于二叉树的层次遍历应用之一
func rightSideView(root *TreeNode) []int {
	if root == nil{
		return []int{}
	}
	var queue []*TreeNode
	var save []int
	queue = append(queue, root)
	for len(queue) > 0 {
		lenth := len(queue)
		//var tmp []int
		save = append(save, queue[lenth-1].Val)
		for i := 0; i < lenth; i ++{
			//tmp = append(tmp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[lenth:]
	}
	return save
}

// 合并区间，排序即可
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	for _, val := range intervals{
		if len(res) == 0 {
			res = append(res, val)
			continue
		}
		if val[0] > res[len(res)-1][1]{
			res = append(res, val)
		} else {
			res[len(res)-1][1] = max(res[len(res)-1][1], val[1])
		}
	}
	return res
}
//func max(x, y int) int{
//	if x < y {
//		return y
//	}
//	return x
//}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 删除排序链表重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := &ListNode{0, head}
	res := cur
	for cur.Next != nil && cur.Next.Next != nil{
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x{
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return res.Next
}

//  下一个排列
func nextPermutation(nums []int)  {
	lenth := len(nums)
	if lenth <= 1{
		return
	}
	//lenth --
	for lenth >= 1 {
		lenth --
		if lenth == 0 || nums[lenth - 1] < nums [lenth]{
			break
		}
	}
	//fmt.Printf("lenth : %d", lenth)
	if lenth == 0 {
		sort.Ints(nums[:])
		return
	}
	lenth --
	for i := len(nums) -1; i > lenth; i -- {
		if nums[i] > nums[lenth] {
			nums[i], nums[lenth] = nums[lenth], nums[i]
			sort.Ints(nums[lenth+1:])
			return
		}
	}
}

// 8 字符串转整数
func myAtoi(s string) int {
	result, sign, i, n := 0, 1, 0, len(s)
	const MinInt32, MaxInt32 = -1 << 31, 1<<31 - 1

	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}

	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}

	for ; i < n; i++ {
		if s[i] < 48 || s[i] > 57 {
			break
		}

		result = result*10 + int(s[i]-'0')
		if sign*result < MinInt32 {
			return MinInt32
		}
		if sign*result > MaxInt32 {
			return MaxInt32
		}
	}

	return sign * result
}

func myAtoiWj(s string) int {
	lenth ,index ,flag := len(s), 0, 1
	for index < lenth {  // 去掉前导字符和处理正负号
		if s[index] == '-' {
			flag = -1
			index ++
			break
		} else if s[index] == '+' {
			index ++
			break
		} else if s[index] >= '0' && s[index] <= '9'{
			break
		} else if s[index] == ' ' {
			index ++
		} else {
			return 0
		}

	}
	res ,max := 0 ,(((1 << 30)-1)<<1) + 1
	min := - max-1
	for index < lenth {
		//fmt.Printf("index : %d,res : %d",index, res)
		if s[index] >= '0' && s[index] <= '9' {
			res = res * 10 + int(s[index]) - 48
			// fmt.Printf("res : %d, int s[index]: %d, s[index] : %v", res, int(s[index]), s[index])
			if flag  * res <= min {
				return min
			} else if flag * res >= max {
				return max
			}
			index ++
		} else {
			break
		}
	}

	return flag * res
}


// 6269. 到目标字符串的最短距离
func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	ans := n
	for i, s := range words {
		if s == target {
			ans = min(ans, min(abs(i-startIndex), n-abs(i-startIndex)))
		}
	}
	if ans == n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	if len1 == 0 || len2 == 0 {
		return 0
	}
	dp := make([][]int, len1 + 1)
	for i := range dp {
		dp[i] = make([]int, len2 + 1)
	}

	for i, v1 := range text1 {
		for j, v2 := range text2 {
			if v1 == v2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len1][len2]
}
//func max (x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}

// 22. 括号生成  递归实现
func generateParenthesis(n int) []string {
	var res []string
	if n < 1 {
		return res
	}
	var digui func(left, right int, oath string)
	digui = func (left, right int, path string) {
		if right == n {
			res = append(res, path)
			return
		}
		if left < n {
			// path += "("
			digui(left+1,right, path + "(")
			// path = path[:len(path)-1]
		}
		if left > right {
			// path += ")"
			digui(left, right + 1, path + ")")
			// path = path[:len(path)-1]
		}
	}
	digui(0,0,"")
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	index := Contains(inorder[:], preorder[0])
	// fmt.Printf("here index : %d\n", index)
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}

func Contains(slice []int, s int) int {
	// fmt.Printf("s : %d\n", s)
	for index, value := range slice {
		// fmt.Printf("index : %d, value : %d\n", index, value)
		if value == s {
			return index
		}
	}
	return len(slice)
}

//

func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre - k]; ok {
			count += m[pre - k]
		}
		m[pre] += 1
		// fmt.Print("m = %v", m)
	}
	return count
}


func main(){
	//head := []int{1,2,3,4,5}
	//nodeHead := &ListNode{Val:0, Next:nil}
	//realHead := nodeHead
	//for val := range head{
	//	node := &ListNode{Val:val, Next:nil}
	//	nodeHead.Next = node
	//	nodeHead = nodeHead.Next
	//}
	//
	//reorderList(realHead.Next)
	//head = []int{}
	//for realHead != nil{
	//	head = append(head, realHead.Val)
	//	realHead = realHead.Next
	//}
	//fmt.Print("head : %v\n", head)
	//
	//fmt.Print("test 22. 括号生成  递归实现 ---------- n =3, result = %v\n",generateParenthesis(3))

	nums := []int{1,1,1}
	target := 2

	fmt.Print(subarraySum(nums, target))
}