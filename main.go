package main

import (
	"fmt"
	"math/rand"
	"time"
)

const boardSize = 5

var rows = [8]int{-1, -1, 0, 1, 1, 1, 0, -1}
var cols = [8]int{0, 1, 1, 1, 0, -1, -1, -1}

/*
	A resource has a terrain name and a count
*/
type Resource struct {
	terrain string
	count   int
}

/*
	Initializes the board with the boundaries and places the desert in the middle.
*/
func initBoard(gameBoard *[boardSize][boardSize]string) {
	// Sets the boundaries of the board for each row.
	(*gameBoard)[0][0] = "X"
	(*gameBoard)[0][4] = "X"
	(*gameBoard)[1][0] = "X"
	(*gameBoard)[3][0] = "X"
	(*gameBoard)[4][0] = "X"
	(*gameBoard)[4][4] = "X"

	// Places the dessert hex in the middle of the board.
	(*gameBoard)[2][2] = "d"
}

/*
	Initializes the resources queue
*/
func initResourceQueue(resouceQueue []Resource) []Resource {
	forest := Resource{
		terrain: "forest",
		count:   4,
	}

	mountains := Resource{
		terrain: "mountain",
		count:   3,
	}

	hills := Resource{
		terrain: "hills",
		count:   3,
	}

	fields := Resource{
		terrain: "fields",
		count:   4,
	}

	pasture := Resource{
		terrain: "pasture",
		count:   4,
	}

	resouceQueue = append(resouceQueue, forest)
	resouceQueue = append(resouceQueue, mountains)
	resouceQueue = append(resouceQueue, hills)
	resouceQueue = append(resouceQueue, fields)
	resouceQueue = append(resouceQueue, pasture)

	return resouceQueue
}

/*
	Shuffles the resources slice to ensure a different solution on each program run.
*/
func shuffleResources(resouceQueue []Resource) []Resource {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(resouceQueue), func(i, j int) { resouceQueue[i], resouceQueue[j] = resouceQueue[j], resouceQueue[i] })
	return resouceQueue
}

/*
	Adds resource to the end of the queue
*/
func enqueue(queue []Resource, resource Resource) []Resource {
	queue = append(queue, resource)
	return queue
}

/*
	Removes the first resource in the queue
*/
func dequeue(queue []Resource) []Resource {
	return queue[1:]
}

/*
	Entry function for backtracking solution.
*/
func createGameBoard(gameBoard *[boardSize][boardSize]string, resouceQueue *[]Resource) {
	startRow := 0
	startCol := 0
	backtrack(gameBoard, *resouceQueue, startRow, startCol)
}

/*
	Using backtracking, the game board is generated.
	Once an answer is found, the search is finalized.
*/
func backtrack(gameBoard *[boardSize][boardSize]string, resouceQueue []Resource, row int, col int) bool {
	if len(resouceQueue) == 0 {
		return true
	}

	if gameBoard[row][col] == "X" || gameBoard[row][col] == "d" {

		newRow, newCol := updateRowAndCol(row, col)

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

			// Ensures that there is still more of that resource remaining before it enqueues it again.
			if currentResource.count != 0 {
				resouceQueue = enqueue(resouceQueue, currentResource)
			}

			if isValidMove(gameBoard, currentResource, row, col) {
				(*gameBoard)[row][col] = currentResource.terrain

				newRow, newCol := updateRowAndCol(row, col)

				foundAnswer := backtrack(gameBoard, resouceQueue, newRow, newCol)

				if foundAnswer {
					return true
				}

				(*gameBoard)[row][col] = "0"
			}

			if currentResource.count == 0 {
				currentResource.count += 1
				resouceQueue = enqueue(resouceQueue, currentResource)
			} else {
				resouceQueue[len(resouceQueue)-1].count += 1
			}
		}
	}

	return false
}

/*
	Helper function to update the row and column.
*/
func updateRowAndCol(row, col int) (int, int) {
	newCol := (col + 1) % boardSize
	newRow := row

	if newCol == 0 {
		newRow += 1
	}

	return newRow, newCol
}

/*
 	Checks if placing a resource at row and column is a valid move.
	All eight locations around a given position is checked.
	The move is valid only if there are no two resources next to each other.
*/
func isValidMove(gameBoard *[boardSize][boardSize]string, currentResource Resource, row int, col int) bool {
	for i := 0; i < len(rows); i++ {
		newRow := row + rows[i]
		newCol := col + cols[i]
		if newRow >= 0 && newRow < boardSize && newCol >= 0 && newCol < boardSize && (*gameBoard)[newRow][newCol] == currentResource.terrain {
			return false
		}
	}

	return true
}

func main() {
	var gameBoard [boardSize][boardSize]string
	var resouceQueue []Resource

	initBoard(&gameBoard)
	resouceQueue = initResourceQueue(resouceQueue)

	resouceQueue = shuffleResources(resouceQueue)

	createGameBoard(&gameBoard, &resouceQueue)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if gameBoard[i][j] != "X" {
				if gameBoard[i][j] == "d" {
					fmt.Print("desert ")
				} else {
					fmt.Printf("%s ", gameBoard[i][j])
				}
			}
		}
		fmt.Printf("\n")
	}
}
