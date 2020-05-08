package main

const boardSize = 5
var rows = [8]int {-1, -1, 0, 1, 1, 1, 0, -1}
var cols = [8]int {0, 1, 1, 1, 0, -1, -1, -1}

type Resource struct {
	terrain 	string
	count		int
	abbrev		string
}

func initBoard(gameBoard *[boardSize][boardSize]rune) {
	// Sets the boundaries of the board for each row.
	(*gameBoard)[0][0] = 'X'
	(*gameBoard)[0][4] = 'X'
	(*gameBoard)[1][0] = 'X'
	(*gameBoard)[3][0] = 'X'
	(*gameBoard)[4][0] = 'X'
	(*gameBoard)[4][4] = 'X'

	// Places the dessert hex in the middle of the board.
	(*gameBoard)[2][2] = 'd'
}

func initResourceQueue(resouceQueue []Resource) []Resource {
	forest := Resource {
		terrain: "forest",
		count: 4,
		abbrev: "f",
	}

	mountains := Resource {
		terrain: "mountain",
		count: 3,
		abbrev: "m",
	}

	hills := Resource {
		terrain: "hills",
		count: 3,
		abbrev: "h",
	}

	fields := Resource {
		terrain: "fields",
		count: 4,
		abbrev: "fi",
	}

	pasture := Resource {
		terrain: "pasture",
		count: 4,
		abbrev: "p",
	}

	resouceQueue = append(resouceQueue, forest)
	resouceQueue = append(resouceQueue, mountains)
	resouceQueue = append(resouceQueue, hills)
	resouceQueue = append(resouceQueue, fields)
	resouceQueue = append(resouceQueue, pasture)

	return resouceQueue
}

func enqueue(queue []Resource, resource Resource) []Resource {
	queue = append(queue, resource)
	return queue
}

func dequeue(queue []Resource) []Resource {
	return queue[1:]
}

func main() {
	var gameBoard [boardSize][boardSize]rune
	var resouceQueue []Resource

	initBoard(&gameBoard)
	initResourceQueue(resouceQueue)
}















