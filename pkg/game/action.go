package game

import "fmt"

func (b *Board) Flag(row, column int) {
	fmt.Println("flag", row, column)
	b.squares[row][column].flag()

	b.Display()
}

func (b *Board) Unflag(row, column int) {
	b.squares[row][column].unflag()
}

func (b *Board) Play(row, column int) {
	fmt.Println("play", row, column)

	if !b.isValid(row, column) {
		fmt.Println("invalid position")
		return
	}

	if b.squares[row][column].isFlagged() {
		fmt.Println("cannot play on flagged square")
		return
	}

	if b.squares[row][column].isBomb {
		b.revealAll()
		fmt.Println("BOOM!")
		return
	}

	b.squares[row][column].reveal()

	if b.squares[row][column].value == 0 || (b.squares[row][column].isRevealed() && b.squares[row][column].value == b.countAdjacentFlag(row, column)) {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				b.playRecursiveUtil(row+i, column+j)
			}
		}
	}

	b.Display()
}

func (b *Board) playRecursiveUtil(row, column int) {
	if !b.isValid(row, column) {
		return
	}

	if b.squares[row][column].isFlagged() || b.squares[row][column].isRevealed() {
		return
	}

	b.squares[row][column].reveal()

	if b.squares[row][column].value != 0 {
		return
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			b.playRecursiveUtil(row+i, column+j)
		}
	}
}

func (b *Board) revealAll() {
	for _, row := range b.squares {
		for _, square := range row {
			square.reveal()
		}
	}

	b.Display()
}

func (b *Board) isValid(row, column int) bool {
	return row >= 0 && row < b.numberRows && column >= 0 && column < b.numberColumns
}

func (b *Board) countAdjacentMines(row, column int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !b.isValid(row+i, column+j) {
				continue
			}

			if b.squares[row+i][column+j].isBomb {
				count++
			}
		}
	}

	return count
}

func (b *Board) countAdjacentFlag(row, column int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !b.isValid(row+i, column+j) {
				continue
			}

			if b.squares[row+i][column+j].isFlagged() {
				count++
			}
		}
	}

	return count
}
