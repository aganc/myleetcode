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



