package main

const (
	BUFFERSIZE    = 1000
	MAZE_ENTRANCE = 'S'
	MAZE_EXIT     = 'E'
	MAZE_WALL     = '#'
	MAZE_PATH     = ' '
	MAZE_TRAIL    = '@'

	MOVE_LEFT  = 0
	MOVE_UP    = 1
	MOVE_RIGHT = 2
	MOVE_DOWN  = 3

	MAZE_NOWAY     = 0
	MAZE_FOUNDEXIT = 1
	MAZE_NOEXIT    = -1

	READY     = 1
	WAITING   = 2
	PROCESSED = 3
)
