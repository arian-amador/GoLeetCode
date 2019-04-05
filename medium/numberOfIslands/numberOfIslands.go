package numberOfIslands

import "fmt"

const (
	land  = '1'
	water = '0'
)

func searchIsland(g [][]byte, row, col, h, w int) {
	g[row][col] = water

	// Upper Boundary
	if row > 0 && g[row-1][col] == land {
		searchIsland(g, row-1, col, h, w)
	}
	// Lower Boundary
	if row+1 < h && g[row+1][col] == land {
		searchIsland(g, row+1, col, h, w)
	}
	// Left Boundary
	if col > 0 && g[row][col-1] == land {
		searchIsland(g, row, col-1, h, w)
	}
	// Right Boundary
	if col+1 < w && g[row][col+1] == land {
		searchIsland(g, row, col+1, h, w)
	}
}

func numIslands(grid [][]byte) int {
	var islands int

	if len(grid) == 0 {
		return islands
	}

	h := len(grid)
	w := len(grid[0])

	for i, rows := range grid {
		for j := range rows {
			curr := grid[i][j]
			if curr == land {
				islands++
				searchIsland(grid, i, j, h, w)
			}
		}
	}

	return islands
}

func debug(grid [][]byte) {
	fmt.Println("-------")
	for i, rows := range grid {
		for j := range rows {
			fmt.Printf("%d ", grid[i][j]-'0')
		}
		fmt.Printf("\n")
	}
	fmt.Println("Height:", len(grid), "Width:", len(grid[0]))
	fmt.Println("-------")
}
