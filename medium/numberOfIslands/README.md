# MEDIUM - Number of Islands [![Build Status](https://api.travis-ci.org/arian-amador/GoLeetCode.svg)](https://travis-ci.org/arian-amador/GoLeetCode)

### Problem

Given a 2d grid map of `1`s (land) and `0`s (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

##### Example
```Go
Input:
11110
11010
11000
00000

Output: 1
```
### Solutions

This is a Depth First Search graph problem making it a recursive search solution.

One of the first things to notice is the function to count islands receives a slices of type `[]byte`. In Go we can use the single quote to convert any character into a byte literal, ex: `'1' == 49`. This was important as we need to check if the current character in the grid is water(0) or land(1).

The solution below is fairly simple as we just iterate over each character starting from the top left and checking if we've found `land`. If `land` is found we send the grid coordinates along with the grid and it's dimensions. We send the grid to the search function allowing us to modify the original grid while continuing the search.

```Go
for i, rows := range grid {
	for j := range rows {
		curr := grid[i][j]
		if curr == land {
			islands++
			searchIsland(&grid, i, j, h, w)
		}
	}
}
```

Prior to calling `searchIsland()` we increment our total island count. This allows us only to track the `land` coordinates that we haven't previously visited during the `searchIsland()` search. Making any `land` found during this loop an `island`.

During our search we check our boundaries so we don't cause an overflow by reading memory outside the grid. At each check we validate if our check will be in bounds and if we've found `land`. If land is found the search function is called with the new coordinates. The search continues until no `land` is found. At this point the function returns to the original caller which could either be a previous `searchIsland()` or the original `numIslands()` function.

```Go
func searchIsland(g *[][]byte, row, col, h, w int) {
	(*g)[row][col] = water

	// Upper Boundry
	if row > 0 && (*g)[row-1][col] == land {
		searchIsland(g, row-1, col, h, w)
	}
	// Lower Boundry
	if row+1 < h && (*g)[row+1][col] == land {
		searchIsland(g, row+1, col, h, w)
	}
	// Left Boundry
	if col > 0 && (*g)[row][col-1] == land {
		searchIsland(g, row, col-1, h, w)
	}
	// Right Boundry
	if col+1 < w && (*g)[row][col+1] == land {
		searchIsland(g, row, col+1, h, w)
	}
}
```

The biggest take away here is that at each `searchIsland()` call we know we're on `land` and by converting the coordinates to `water` we're marking them as visited. This is also why we pass `grid` to the `searchIsland()` function. It allows us to modify the original grid as we search.

Go passes slices a bit different as it sends the header address of the slice to the receiving function as a pointer. Meaning that any changes to existing slice elements in the receiving function will modify the original slice. If the slice is appended or an element removed then a new copy is made.

##### Additional Resources

- [Recursion](https://www.geeksforgeeks.org/recursion/)
- [Depth First Search or DFS for a Graph](https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/)
- [Slices: The slice header](https://blog.golang.org/slices)

---

#### Hope you find this useful!
