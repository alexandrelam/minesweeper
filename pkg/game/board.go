package game

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

type Board struct {
	numberRows     int
	numberColumns  int
	numberBombs    int
	numberDugTiles int
	squares        [][]*Square
}

type IBoard interface {
	NewBoard(numberRows, numberColumns, numberBombs, numberDugTiles int) *Board // create a new board
	Play(row, column int)                                                       // Play a square
	Flag(row, column int)                                                       // Flag a square
	Unflag(row, column int)                                                     // Unflag a square
	Display()                                                                   // display board with hidden squares
	DisplayNoHidden()                                                           // display board without hidden squares
}

func NewBoard(numberRows, numberColumns, numberBombs int) *Board {

	if numberBombs > numberRows*numberColumns {
		panic("too many bombs")
	}

	// init board
	b := &Board{
		numberRows:     numberRows,
		numberColumns:  numberColumns,
		numberBombs:    numberBombs,
		numberDugTiles: 0,
		squares:        make([][]*Square, numberRows),
	}

	for i := 0; i < numberRows; i++ {
		for j := 0; j < numberColumns; j++ {
			b.squares[i] = append(b.squares[i], newSquare(false))
		}
	}

	// place bombs
	numberOfBombs := 0
	for numberOfBombs < numberBombs {
		row := rand.Intn(numberRows)
		column := rand.Intn(numberColumns)

		if b.squares[row][column].IsBomb {
			continue
		}

		b.squares[row][column].IsBomb = true
		numberOfBombs++
	}

	// calculate values
	for row := 0; row < numberRows; row++ {
		for column := 0; column < numberColumns; column++ {
			if b.squares[row][column].IsBomb {
				continue
			}

			b.squares[row][column].Value = b.countAdjacentMines(row, column)

		}
	}

	return b
}

func (b *Board) Display() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Bold)
	redlte := color.New(color.FgRed)

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
				if b.squares[row][column].IsBomb {
					print(" B ")
				} else {
					if b.squares[row][column].Value == 0 {
						print("   ")
					} else {
						fmt.Printf(" %d ", b.squares[row][column].Value)
					}
				}
			} else if b.squares[row][column].isFlagged() {
				redlte.Print(" F ")
			} else {
				print(" . ")
			}
		}
		println()
	}
	println()
}

func (b *Board) DisplayNoHidden() {
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

			if b.squares[row][column].IsBomb {
				print(" B ")
			} else {
				if b.squares[row][column].Value == 0 {
					print("   ")
				} else {
					fmt.Printf(" %d ", b.squares[row][column].Value)
				}
			}
		}
		println()
	}
	println()
}

func (b *Board) GetSquare() [][]*Square {
	return b.squares
}
