package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must specify the filename of your maze")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("Too many command line arguments")
		os.Exit(1)
	}

	filename := os.Args[1]
	maze := Maze{}

	GetMazeFromFile(filename, &maze)

	// Success on reading the array
	fmt.Printf("Start Found: %d, %d\n", maze.startX, maze.startY)
	fmt.Printf("Exit Found: %d, %d\n", maze.exitX, maze.exitY)
	fmt.Printf("Grid Size: %d, %d\n", maze.numRows, maze.numCols)

	// Uncomment and implement solve function if needed
	// if solve(&maze) == MAZE_FOUNDEXIT {
	//     fmt.Println("Found exit!")
	// } else {
	//     fmt.Println("Can't reach exit!")
	// }

	PrintMaze(&maze)
	findPath(&maze)
	PrintMaze(&maze)
}
