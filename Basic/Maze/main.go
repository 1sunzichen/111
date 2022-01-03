package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int
	//扫描前2个数：行数和列数，继续执行这个方法 扫描后面的字符
	//https://studygolang.com/pkgdoc
	fmt.Fscanf(file, "%d %d", &row, &col)
	//
	fmt.Println(row, col)
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

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(m [][]int, start, end point) [][]int {
	//路径
	step := make([][]int, len(m))
	for i := range step {
		step[i] = make([]int, len(m[i]))

	}
	Q := []point{start}
	//队列不空的时候 走下去
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)

			//maze at next is 0
			//and steps at n
			val, ok := next.at(m)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(step)
			if !ok || val == 1 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(step)
			step[next.i][next.j] = curSteps + 1
			Q = append(Q, next)

		}

	}
	return step
}

func main() {
	maze := readMaze("Basic/Maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}
	//         起始点              结束点 坐标 行数-1,列数-1
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
