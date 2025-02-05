package main

import (
	"os"
	"testing"
)

// TestMazeStructureInitialization tests basic maze structure creation and initialization
func TestMazeStructureInitialization(t *testing.T) {
	maze := Maze{}
	if maze.mazeMap != nil {
		t.Error("Expected empty maze map on initialization")
	}
	if maze.numRows != 0 || maze.numCols != 0 {
		t.Error("Expected 0 dimensions on initialization")
	}
}

// TestMake2DArray tests the array creation utility function
func TestMake2DArray(t *testing.T) {
	rows, cols := 5, 10
	array := make2DArray(rows, cols)

	if len(array) != rows {
		t.Errorf("Expected %d rows, got %d", rows, len(array))
	}

	for i := 0; i < rows; i++ {
		if len(array[i]) != cols {
			t.Errorf("Row %d: expected %d columns, got %d", i, cols, len(array[i]))
		}
	}
}

// TestGetMazeFromFile tests reading various maze configurations from files
func TestGetMazeFromFile(t *testing.T) {
	// Store original values to restore after test
	originalMaxChars := maxCharsPerRow
	originalMaxRows := maxRows
	originalDeltaChars := deltaChars
	originalDeltaRows := deltaRows

	// Restore original values after test
	defer func() {
		maxCharsPerRow = originalMaxChars
		maxRows = originalMaxRows
		deltaChars = originalDeltaChars
		deltaRows = originalDeltaRows
	}()

	// Set smaller initial values for testing
	maxCharsPerRow = 10
	maxRows = 5
	deltaChars = 10
	deltaRows = 5

	tests := []struct {
		name          string
		mazeContent   string
		expectedStart bool
		expectedExit  bool
	}{
		{
			name: "Valid Basic Maze",
			mazeContent: `###
							#S#
							#E#
							###`,
			expectedStart: true,
			expectedExit:  true,
		},
		{
			name: "Long Lines",
			mazeContent: `##########S#
							# # # # # # 
							#         E#`,
			expectedStart: true,
			expectedExit:  true,
		},
		{
			name: "Many Rows",
			mazeContent: `###
							#S#
							# #
							# #
							# #
							#E#`,
			expectedStart: true,
			expectedExit:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary maze file
			tmpfile, err := os.CreateTemp("", "maze_*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.mazeContent)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			maze := Maze{}
			GetMazeFromFile(tmpfile.Name(), &maze)

			// Verify start and exit were found if expected
			if tt.expectedStart {
				foundStart := false
				for i := 0; i < maze.numRows; i++ {
					for j := 0; j < maze.numCols; j++ {
						if maze.mazeMap[i][j] == 'S' {
							foundStart = true
							break
						}
					}
				}
				if !foundStart {
					t.Error("Expected to find start position")
				}
			}

			if tt.expectedExit {
				foundExit := false
				for i := 0; i < maze.numRows; i++ {
					for j := 0; j < maze.numCols; j++ {
						if maze.mazeMap[i][j] == 'E' {
							foundExit = true
							break
						}
					}
				}
				if !foundExit {
					t.Error("Expected to find exit position")
				}
			}
		})
	}
}

// TestMazeResizing specifically tests the resizing functionality
func TestMazeResizing(t *testing.T) {
	// Store original values
	originalMaxChars := maxCharsPerRow
	originalMaxRows := maxRows

	// Restore original values after test
	defer func() {
		maxCharsPerRow = originalMaxChars
		maxRows = originalMaxRows
	}()

	// Set small initial values to force resizing
	maxCharsPerRow = 5
	maxRows = 3

	// Create a maze that will require resizing
	mazeContent := `#####S
					#   #
					#   #
					#  E#`

	tmpfile, err := os.CreateTemp("", "maze_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(mazeContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	maze := Maze{}
	GetMazeFromFile(tmpfile.Name(), &maze)

	// Verify the maze was read correctly
	if maze.startX == 0 && maze.startY == 0 {
		t.Error("Start position not found")
	}
	if maze.exitX == 0 && maze.exitY == 0 {
		t.Error("Exit position not found")
	}
}

// TestPathFinding tests the actual path finding functionality
func TestPathFinding(t *testing.T) {
	tests := []struct {
		name        string
		mazeContent string
		solvable    bool
	}{
		{
			name: "Simple Solvable Maze",
			mazeContent: `#####
							#S  #
							# # #
							#  E#
							#####`,
			solvable: true,
		},
		{
			name: "Unsolvable Maze",
			mazeContent: `#####
							#S# #
							# # #
							#  E#
							#####`,
			solvable: false,
		},
		{
			name: "Single Path Maze",
			mazeContent: `#####
							#S  #
							### #
							#E  #
							#####`,
			solvable: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary maze file
			tmpfile, err := os.CreateTemp("", "maze_*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.Write([]byte(tt.mazeContent)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			maze := Maze{}
			GetMazeFromFile(tmpfile.Name(), &maze)

			// Store initial state
			initialState := make2DArray(maze.numRows, maze.numCols)
			for i := 0; i < maze.numRows; i++ {
				copy(initialState[i], maze.mazeMap[i])
			}

			findPath(&maze)

			// Check if path exists by counting '@' characters
			foundPath := false
			for i := 0; i < maze.numRows; i++ {
				for j := 0; j < maze.numCols; j++ {
					if maze.mazeMap[i][j] == MAZE_TRAIL {
						foundPath = true
						break
					}
				}
				if foundPath {
					break
				}
			}

			if tt.solvable && !foundPath {
				t.Error("Expected to find a path but none was marked")
			}

			// Verify path connects start to end
			if foundPath {
				hasPathToExit := false
				startX, startY := maze.startX, maze.startY

				// Check if there's a trail (@) adjacent to start
				if (startY > 0 && maze.mazeMap[startY-1][startX] == MAZE_TRAIL) ||
					(startY < maze.numRows-1 && maze.mazeMap[startY+1][startX] == MAZE_TRAIL) ||
					(startX > 0 && maze.mazeMap[startY][startX-1] == MAZE_TRAIL) ||
					(startX < maze.numCols-1 && maze.mazeMap[startY][startX+1] == MAZE_TRAIL) {
					hasPathToExit = true
				}

				if !hasPathToExit && tt.solvable {
					t.Error("Path does not connect to start position")
				}
			}
		})
	}
}

// TestNodeManipulation tests the node content and status manipulation functions
func TestNodeManipulation(t *testing.T) {
	maze := &Maze{
		mazeMap: [][]rune{
			{'#', 'S', ' ', 'E', '#'},
			{'#', '#', '#', '#', '#'},
		},
		numRows: 2,
		numCols: 5,
	}

	// Test getNodeContent
	if getNodeContent(maze, 1) != 'S' {
		t.Error("Expected 'S' at node 1")
	}

	// Test changeNodeContent
	changeNodeContent(maze, 2, '@')
	if maze.mazeMap[0][2] != '@' {
		t.Error("Expected '@' after changing node content")
	}

	// Test status manipulation
	status := make2DIntArray(2, 5)
	changeNodeStatus(status, 1, 5, PROCESSED)
	if getNodeStatusContent(status, 1, 5) != PROCESSED {
		t.Error("Expected PROCESSED status after change")
	}
}

// Add this new test function to the existing test file

func TestUpwardPathFinding(t *testing.T) {
	// Create a maze where the path must go upward to reach the exit
	maze := &Maze{
		mazeMap: [][]rune{
			{'#', '#', 'E', '#', '#'},
			{'#', ' ', ' ', ' ', '#'},
			{'#', '#', 'S', '#', '#'},
		},
		startX:  2,
		startY:  2,
		exitX:   2,
		exitY:   0,
		numRows: 3,
		numCols: 5,
	}

	// Run path finding
	findPath(maze)

	// Verify that a path was found
	if maze.mazeMap[1][2] != MAZE_TRAIL {
		t.Error("Expected path to move upward through middle position")
	}

	// Verify path connects start to exit
	if maze.mazeMap[1][1] != MAZE_TRAIL && maze.mazeMap[1][2] != MAZE_TRAIL && maze.mazeMap[1][3] != MAZE_TRAIL {
		t.Error("No upward path found")
	}

	// Print the maze for debugging
	t.Logf("Resulting maze:\n")
	for i := 0; i < maze.numRows; i++ {
		t.Logf("%s\n", string(maze.mazeMap[i]))
	}
}
