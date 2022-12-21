package game

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

type Board struct {
	numberRows    int
	numberColumns int
	numberBombs   int
	squares       [][]*Square
}

func NewBoard(numberRows, numberColumns, numberBombs int) *Board {

	if numberBombs > numberRows*numberColumns {
		panic("too many bombs")
	}

	// init board
	b := &Board{
		numberRows:    numberRows,
		numberColumns: numberColumns,
		squares:       make([][]*Square, numberRows),
	}

	for i := 0; i < numberRows; i++ {
		for j := 0; j < numberColumns; j++ {
			b.squares[i] = append(b.squares[i], newSquare(false))
		}
	}

	// place bombs
	for i := 0; i < numberBombs; i++ {
		row := rand.Intn(numberRows)
		column := rand.Intn(numberColumns)

		b.squares[row][column].bomb()
	}

	// calculate values
	for row := 0; row < numberRows; row++ {
		for column := 0; column < numberColumns; column++ {
			if b.squares[row][column].isBomb {
				continue
			}

			b.squares[row][column].value = b.countAdjacentMines(row, column)

		}
	}

	return b
}

func (b *Board) Display() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Bold)

	fmt.Printf("   ")
	for row := 0; row < b.numberRows; row++ {
		boldWhite.Printf("%2d ", row)
	}
	println()

	for row := 0; row < b.numberRows; row++ {
		for column := 0; column < b.numberColumns; column++ {
			if column == 0 {
				boldWhite.Printf("%2d ", row)
			}

			if b.squares[row][column].isRevealed() {
				if b.squares[row][column].isBomb {
					print(" B ")
				} else {
					fmt.Printf(" %d ", b.squares[row][column].value)
				}
			} else if b.squares[row][column].isFlagged() {
				print(" F ")
			} else {
				print(" . ")
			}
		}
		println()
	}
	println()
}

func (b *Board) DisplayNoHidden() {
	for row := 0; row < b.numberRows; row++ {
		for column := 0; column < b.numberColumns; column++ {
			if b.squares[row][column].isBomb {
				print("B")
			} else if b.squares[row][column].isFlagged() {
				print("F")
			} else {
				print(b.squares[row][column].value)
			}
		}
		println()
	}
}
