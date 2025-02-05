# MazeRunnerGo

MazeRunnerGo is a Go-based maze solver that finds the shortest path from a start position to an exit in a maze. 
The program reads a maze from a text file and prints the solution in the console.

## Features
Reads a maze from a text file.

Uses pathfinding algorithms to find the shortest route.

Prints the maze in the terminal.

Utilizes recursion and backtracking to solve the maze.

## Prerequisites
Ensure you have Go installed on your system. You can download it from:
[https://go.dev/dl/](https://go.dev/dl/)

## Installation
1. Clone this repository:
    ```bash
    git clone https://github.com/yourusername/MazeRunnerGo.git
    cd MazeRunnerGo
    ```
2. Ensure you have the required dependencies installed.
3. Place your maze file (e.g., maze1.txt) in the project directory. The maze should be in the following format:

```plaintext
###################
#       #         #
#   #  E#         #
#   #####         #
#                 #
############## ####
#   S             #
###################
```

Where:
- `#` represents walls.
- S is the starting position.
- E is the exit.
- Spaces ( ) are paths.
- Traversed path is marked with @.

## Running the Program
Compile and run the program using:
```bash
go run main.go
```

This will output the maze in the terminal with the first print being the initial maze and the second one being the solved version.

Example Output:
```plaintext
./maze_solver maze.txt
Start Found: 4, 6
Exit Found: 7, 2
Grid Size: 8, 50
###################
#       #         #
#   #  E#         #
#   #####         #
#                 #
############## ####
#   S             #
###################
###################
#  @@@@@#         #
#  @#  @#         #
#  @#####         #
#  @@@@@@@@@@@@   #
##############@####
#   @@@@@@@@@@@   #
###################
```

## How It Works
1. Loads the maze from a text file.
2. Identifies the start (S) and exit (E) positions.
3. Uses a recursive depth-first search (DFS) to find the shortest path.
4. Animates the traversal by marking the visited path with @.
5. Displays the solution dynamically in the terminal.
