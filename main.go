package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize = 5
	numberOfResources = 5
)

var rows = [8]int {-1, -1, 0, 1, 1, 1, 0, -1}
var cols = [8]int {0, 1, 1, 1, 0, -1, -1, -1}

type Resource struct {
	terrain 	string
	count		int
	abbrev		rune
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
		abbrev: 'f',
	}

	mountains := Resource {
		terrain: "mountain",
		count: 3,
		abbrev: 'm',
	}

	hills := Resource {
		terrain: "hills",
		count: 3,
		abbrev: 'h',
	}

	fields := Resource {
		terrain: "fields",
		count: 4,
		abbrev: 'i',
	}

	pasture := Resource {
		terrain: "pasture",
		count: 4,
		abbrev: 'p',
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

func createGameBoard(gameBoard *[boardSize][boardSize]rune, resouceQueue *[]Resource) {
	startRow := 0
	startCol := 0
	result := backtrack(gameBoard, *resouceQueue, startRow, startCol)

	fmt.Println(result)
}

func backtrack(gameBoard *[boardSize][boardSize]rune, resouceQueue []Resource, row int, col int) bool {
	if len(resouceQueue) == 0 {
		return true
	}

	if gameBoard[row][col] == 'X' || gameBoard[row][col] == 'd' {
		newCol := (col + 1) % boardSize
		newRow := row

		if newCol == 0 {
			newRow += 1
		}

		foundAnswer := backtrack(gameBoard, resouceQueue, newRow, newCol)

		if foundAnswer {
			return true
		}
	} else {
		queueLength := len(resouceQueue)
		for i := 0; i < queueLength; i++ {
			currentResource := (resouceQueue)[0]

			resouceQueue = dequeue(resouceQueue)

			currentResource.count -= 1

			// Ensures that there is more of the resource left before it enqueues it.
			if currentResource.count != 0 {
				resouceQueue = enqueue(resouceQueue, currentResource)
			}

			if isValidMove(gameBoard, currentResource, row, col) {
				(*gameBoard)[row][col] = currentResource.abbrev

				newCol := (col + 1) % boardSize
				newRow := row

				if newCol == 0 {
					newRow += 1
				}

				foundAnswer := backtrack(gameBoard, resouceQueue, newRow, newCol)

				if foundAnswer {
					return true
				}

				(*gameBoard)[row][col] = '0'
			}

			if currentResource.count == 0 {
				currentResource.count += 1
				resouceQueue = enqueue(resouceQueue, currentResource)
			} else {
				resouceQueue[len(resouceQueue) - 1].count += 1
			}
		}
	}

	return false
}

func isValidMove(gameBoard *[boardSize][boardSize]rune, currentResource Resource, row int, col int) bool {
	for i := 0; i < len(rows); i++ {
		newRow := row + rows[i]
		newCol := col + cols[i]
		if newRow >= 0 && newRow < boardSize && newCol >= 0 && newCol < boardSize && (*gameBoard)[newRow][newCol] == currentResource.abbrev {
			return false
		}
	}

	return true
}

func shuffleResources(resouceQueue []Resource) []Resource {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(resouceQueue), func(i, j int) {resouceQueue[i], resouceQueue[j] = resouceQueue[j], resouceQueue[i]})
	return resouceQueue
}

func main() {
	var gameBoard [boardSize][boardSize]rune
	var resouceQueue []Resource

	initBoard(&gameBoard)
	resouceQueue = initResourceQueue(resouceQueue)

	resouceQueue = shuffleResources(resouceQueue)

	createGameBoard(&gameBoard, &resouceQueue)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Printf("%c ", gameBoard[i][j])
		}
		fmt.Printf("\n")
	}
}















