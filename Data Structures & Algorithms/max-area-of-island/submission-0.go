func maxAreaOfIsland(grid [][]int) int {
	size := 0
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		//check out of bounds
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return 0	
		}
		curCell := grid[row][col]
		//check if water or already visited. If yes we stop searching
		if curCell == 0 || visited[row][col]{
			return 0
		}
		// set visited for next recursion
		visited[row][col] = true
		return 1 + dfs(row-1,col) + dfs(row+1,col) + dfs(row,col-1) + dfs(row,col+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				local := dfs(i,j)
				if local > size {
					size = local
				}				
			}
		}
	}
	return size
}
