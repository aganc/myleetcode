package echo


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