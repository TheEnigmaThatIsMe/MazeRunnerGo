package main

import (
	"bufio"
	"fmt"
	"os"
)

type Maze struct {
	mazeMap [][]rune
	startX  int
	startY  int
	exitX   int
	exitY   int
	numRows int
	numCols int
	initDir int
}

var maxCharsPerRow = 50
var maxRows = 10
var deltaChars = 10
var deltaRows = 5

func make2DArray(rows int, columns int) [][]rune {
	array := make([][]rune, rows)
	for i := 0; i < rows; i++ {
		array[i] = make([]rune, columns)
	}
	return array
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetMazeFromFile(filename string, maze *Maze) {
	success := 0
	numberOfRows := 0
	foundStart := 0
	foundExit := 0
	iLoop := 0

	var ioMaze [][]rune
	currentMaxChars := maxCharsPerRow

	for success == 0 {
		// Initialize start and exit positions
		foundStart = 0
		foundExit = 0

		success = 1 // assume that the maze is read in correctly, if not, set success to 0 so it tries again.

		file, err := os.Open(filename)
		check(err)
		defer func(file *os.File) {
			err := file.Close()
			check(err)
		}(file)

		// Allocate maze size
		ioMaze = make2DArray(maxRows, currentMaxChars)

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ioLine := scanner.Text()
			length := len(ioLine)

			if length >= currentMaxChars {
				// It is possible that we have a string longer than the currentMaxChars size
				// The file needs to be reset with longer line lengths
				_, err := file.Seek(0, 0)
				check(err)
				currentMaxChars += deltaChars
				numberOfRows = 0
				ioMaze = make2DArray(maxRows, currentMaxChars)
				success = 0
				break
			}

			if numberOfRows >= maxRows {
				// It is possible that we have a string longer than the currentMaxChars size
				// The file needs to be reset with longer line lengths
				_, err := file.Seek(0, 0)
				check(err)
				maxRows += deltaRows
				numberOfRows = 0
				ioMaze = make2DArray(maxRows, currentMaxChars)
				success = 0
				break
			}

			// The line we just read in falls within our grid array size, so add it to the grid
			for i := 0; i < len(ioLine); i++ {
				ioMaze[numberOfRows][i] = rune(ioLine[i])
			}

			// check for start / exit in maze
			if foundStart == 0 {
				for iLoop = 0; iLoop < len(ioLine); iLoop++ {
					if ioMaze[numberOfRows][iLoop] == 'S' {
						maze.startX = iLoop
						maze.startY = numberOfRows
						foundStart = 1
					}
				}
			}

			if foundExit == 0 {
				for iLoop = 0; iLoop < len(ioLine); iLoop++ {
					if ioMaze[numberOfRows][iLoop] == 'E' {
						maze.exitX = iLoop
						maze.exitY = numberOfRows
						foundExit = 1
					}
				}
			}

			numberOfRows++
		}

		check(scanner.Err())
	}

	// rows should return number of lines read into the array
	maze.numRows = numberOfRows
	maze.numCols = currentMaxChars
	maze.mazeMap = ioMaze
}

func PrintMaze(maze *Maze) {
	for n := 0; n < maze.numRows; n++ {
		fmt.Println(string(maze.mazeMap[n]))
	}
}
