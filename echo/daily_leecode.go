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


