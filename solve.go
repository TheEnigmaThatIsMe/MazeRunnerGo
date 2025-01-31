package main

func make2DIntArray(rows int, columns int) [][]int {
	array := make([][]int, rows)
	for i := range array {
		array[i] = make([]int, columns)
		for j := range array[i] {
			array[i][j] = READY
		}
	}
	return array
}

func getNodeContent(maze *Maze, iNodeNumber int) rune {
	nRow := iNodeNumber / maze.numCols
	nCol := iNodeNumber - (nRow * maze.numCols)
	if maze.mazeMap[nRow][nCol] == MAZE_EXIT {
		return MAZE_PATH
	}
	return maze.mazeMap[nRow][nCol]
}

func changeNodeContent(maze *Maze, iNodeNumber int, pathChar rune) {
	nRow := iNodeNumber / maze.numCols
	nCol := iNodeNumber - (nRow * maze.numCols)
	maze.mazeMap[nRow][nCol] = pathChar
}

func getNodeStatusContent(mazeStatus [][]int, iNodeNumber int, cols int) int {
	nRow := iNodeNumber / cols
	nCol := iNodeNumber - (nRow * cols)
	return mazeStatus[nRow][nCol]
}

func changeNodeStatus(mazeStatus [][]int, iNodeNumber int, cols int, status int) {
	nRow := iNodeNumber / cols
	nCol := iNodeNumber - (nRow * cols)
	mazeStatus[nRow][nCol] = status
}

func shortestPath(maze *Maze, nodeStart int, nodeExit int) {
	rows := maze.numRows
	cols := maze.numCols
	maxSize := rows * cols
	queue := make([]int, maxSize)
	origin := make([]int, maxSize)
	i := 0

	for i := 0; i < maxSize; i++ {
		queue[i] = 0
		origin[i] = 0
	}

	front := 0
	rear := 0

	// create array to hold the status of each grid point we have checked
	mazeStatus := make2DIntArray(rows, cols)

	queue[rear] = nodeStart
	origin[rear] = -1
	rear++

	var current, left, right, top, down int

	for front != rear { // while the queue is not empty
		if current == 157 {
			left = current
		}

		if queue[front] == nodeExit {
			break // maze is solved
		}

		current = queue[front]
		left = current - 1

		if left >= 0 && left/cols == current/cols { //if left node exists
			if getNodeContent(maze, left) == MAZE_PATH {
				if getNodeStatusContent(mazeStatus, left, cols) == READY {
					queue[rear] = left
					origin[rear] = current
					changeNodeStatus(mazeStatus, left, cols, WAITING)
					rear++
				}
			}
		}

		right = current + 1
		if right < maxSize && (right/cols) == (current/cols) {
			if getNodeContent(maze, right) == MAZE_PATH {
				if getNodeStatusContent(mazeStatus, right, cols) == READY {
					queue[rear] = right
					origin[rear] = current
					changeNodeStatus(mazeStatus, right, cols, WAITING)
					rear++
				}
			}
		}

		top = current - cols
		if top >= 0 && top < maxSize {
			if getNodeContent(maze, top) == MAZE_PATH {
				if getNodeStatusContent(mazeStatus, top, cols) == READY {
					queue[rear] = top
					origin[rear] = current
					changeNodeStatus(mazeStatus, top, cols, WAITING)
					rear++
				}
			}
		}

		down = current + cols
		if down < maxSize && down > 0 {
			if getNodeContent(maze, down) == MAZE_PATH {
				if getNodeStatusContent(mazeStatus, down, cols) == READY {
					queue[rear] = down
					origin[rear] = current
					changeNodeStatus(mazeStatus, down, cols, WAITING)
					rear++
				}
			}
		}

		changeNodeStatus(mazeStatus, current, cols, PROCESSED)
		front++

	}

	// Update the maze with the path
	current = nodeExit
	changeNodeContent(maze, nodeExit, MAZE_TRAIL)
	for i = front; i >= 0; i-- {
		if queue[i] == current {
			current = origin[i]
			if current == -1 {
				return // maze is solved
			}
			changeNodeContent(maze, current, MAZE_TRAIL)
		}
	}
}

func findPath(theMaze *Maze) {
	iStartNode := theMaze.startY*theMaze.numCols + theMaze.startX
	iExitNode := theMaze.exitY*theMaze.numCols + theMaze.exitX

	shortestPath(theMaze, iStartNode, iExitNode)
}
