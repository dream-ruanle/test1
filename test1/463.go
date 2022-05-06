package main

import "fmt"

/*
463. 岛屿的周长
给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。

网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。
 */
type direction struct {
	x,y int
}
var directions =[]direction{{0,1},{0,-1},{1,0},{-1,0}}

func islandPerimeter(grid [][]int) int {

	m,n:=len(grid),len(grid[0])
	var dfs463 func(int,int)int
	dfs463 = func(i, j int) int{

		if  i>=m||i<0||j>=n||j<0||grid[i][j]==0 {return 1}
		if grid[i][j]==-1{return 0}
		grid[i][j]=-1
		return dfs463(i+1,j)+dfs463(i-1,j)+dfs463(i,j+1)+dfs463(i,j-1)
	}

	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]==1{
				return dfs463(i,j)


			}
		}
	}
	return 0
}

func main() {
	test:=[][]int{{0,1,0,0},{1,1,1,0},{0,1,0,0},{1,1,0,0}}
	fmt.Print(islandPerimeter(test))
}
