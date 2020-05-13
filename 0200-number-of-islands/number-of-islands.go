package _200_number_of_islands

func zeroIsland(grid *[][]byte, i, j int) {
	if i == len(*grid) || j == len((*grid)[i]) {
		return
	}
	if (*grid)[i][j] == 0 || (*grid)[i][j] == '0' {
		return
	}
	(*grid)[i][j] = 0
	if i > 0 {
		zeroIsland(grid, i-1, j)
	}
	if i < len(*grid)-1 {
		zeroIsland(grid, i+1, j)
	}
	if j > 0 {
		zeroIsland(grid, i, j-1)
	}
	if j < len((*grid)[i])-1 {
		zeroIsland(grid, i, j+1)
	}
}

func numIslands(grid [][]byte) (islands int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 || grid[i][j] == '1' {
				islands++
				zeroIsland(&grid, i, j)
			}
		}
	}
	return
}
