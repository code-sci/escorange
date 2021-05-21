package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//
// 创建人: huangjinmu
// 创建时间: 2021/5/21 上午10:11
func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		println(rand.Int31n(1000000))
	}
}

/*func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, col := range row {
			fmt.Printf("%3d ", col)
		}
		fmt.Println()
	}
	fmt.Println("----------------")
	aa := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range aa {
		for _, val := range row {
			fmt.Printf("%4d", val)
		}
		fmt.Println()
	}
}*/

/**
	i i i i i
j	0 1 0 0 0
j	0 0 0 1 0
j	0 1 0 1 0
j	1 1 1 0 0
j	0 1 0 0 1
j	0 1 0 0 0
*/
var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) at(grid [][]int) (int, bool) {
	// i 越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// J 越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := []point{start} // Q = 起点{0, 0}
	for len(Q) > 0 {    // len(Q) == 1, 此处 = true
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		// 开始朝四个方向搜索, 可以走过去的 + 1
		for _, dir := range dirs {
			next := cur.add(dir)

			val, ok := next.at(maze) // 是否越界
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}
