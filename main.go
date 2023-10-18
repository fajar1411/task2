package main

import (
	"fmt"
	"math/rand"
)

const (
	obstacle   = "#"
	clearPath  = "."
	player     = "X"
	hiddenItem = "$"
	rows       = 6
	columns    = 8
)

type HiddenItemGame struct {
	grid             [][]string
	playerPos        struct{ row, col int }
	hiddenItemPos    struct{ row, col int }
	probableItemPos  []struct{ row, col int }
	steps            struct{ up, right, down int }
	probableItemList []struct{ row, col int }
}

func NewHiddenItemGame() *HiddenItemGame {
	game := &HiddenItemGame{
		grid: make([][]string, rows),
		playerPos: struct{ row, col int }{
			row: 4,
			col: 1,
		},
		steps: struct{ up, right, down int }{
			up:    1,
			right: 2,
			down:  3,
		},
	}

	for i := range game.grid {
		game.grid[i] = make([]string, columns)
		for j := range game.grid[i] {
			game.grid[i][j] = clearPath
		}
	}

	game.grid[game.playerPos.row][game.playerPos.col] = player

	game.placeHiddenItem()

	game.calculateProbableItemLocations()

	return game
}

func (game *HiddenItemGame) placeHiddenItem() {

	for {
		row := rand.Intn(rows)
		col := rand.Intn(columns)

		if game.grid[row][col] == clearPath {
			game.hiddenItemPos = struct{ row, col int }{row, col}
			break
		}
	}
}

func (game *HiddenItemGame) calculateProbableItemLocations() {
	game.probableItemPos = append(game.probableItemPos, struct{ row, col int }{
		row: game.playerPos.row - game.steps.up,
		col: game.playerPos.col,
	})
	game.probableItemPos = append(game.probableItemPos, struct{ row, col int }{
		row: game.playerPos.row,
		col: game.playerPos.col + game.steps.right,
	})
	game.probableItemPos = append(game.probableItemPos, struct{ row, col int }{
		row: game.playerPos.row + game.steps.down,
		col: game.playerPos.col,
	})

	for _, pos := range game.probableItemPos {
		if pos.row >= 0 && pos.row < rows && pos.col >= 0 && pos.col < columns && game.grid[pos.row][pos.col] == clearPath {
			game.probableItemList = append(game.probableItemList, pos)
		}
	}
}

func (game *HiddenItemGame) displayGrid() {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == game.playerPos.row && j == game.playerPos.col {
				fmt.Print(player, " ")
			} else if i == game.hiddenItemPos.row && j == game.hiddenItemPos.col {
				fmt.Print(hiddenItem, " ")
			} else if contains(game.probableItemList, struct{ row, col int }{i, j}) {
				fmt.Print(hiddenItem, " ")
			} else {
				fmt.Print(game.grid[i][j], " ")
			}
		}
		fmt.Println()
	}
}

func contains(slice []struct{ row, col int }, element struct{ row, col int }) bool {
	for _, v := range slice {
		if v.row == element.row && v.col == element.col {
			return true
		}
	}
	return false
}

func main() {
	game := NewHiddenItemGame()

	fmt.Println("Grid:")
	game.displayGrid()

	fmt.Println("\nProbable Item Locations:")
	fmt.Println(game.probableItemList)
}
