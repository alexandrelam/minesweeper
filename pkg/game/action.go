package game

import "fmt"

type PlayReturn struct {
	IsPlayed bool
	IsWin    bool
	IsLost   bool
}

// return: true if the square is valid, false otherwise
func (b *Board) Flag(row, column int) bool {
	if !b.isValid(row, column) {
		return false
	}

	if b.squares[row][column].isRevealed() {
		return false
	}

	if b.squares[row][column].isFlagged() {
		return false
	}

	b.squares[row][column].flag()
	return true
}

// return: true if the square is valid, false otherwise
func (b *Board) Unflag(row, column int) bool {
	if !b.isValid(row, column) {
		return false
	}

	if b.squares[row][column].isRevealed() {
		return false
	}

	if !b.squares[row][column].isFlagged() {
		return false
	}

	b.squares[row][column].unflag()
	return true
}

func (b *Board) Play(row, column int) PlayReturn {
	if !b.isValid(row, column) {
		fmt.Println("invalid position")
		return PlayReturn{false, false, false}
	}

	if b.squares[row][column].isFlagged() {
		fmt.Println("cannot play on flagged square")
		return PlayReturn{false, false, false}
	}

	if b.squares[row][column].IsBomb {
		b.revealAll()
		return PlayReturn{true, false, true}
	}

	if !b.squares[row][column].isRevealed() {
		b.squares[row][column].reveal()
		b.numberDugTiles++
	}

	// if tile is doesn't have bomb around it, bomb cannot explode
	if b.squares[row][column].Value == 0 {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				hasExploded := b.playRecursiveUtil(row+i, column+j, false)
				if hasExploded {
					b.revealAll()
					return PlayReturn{true, false, true}
				}
			}
		}

	}

	// else bomb can explode
	if b.squares[row][column].isRevealed() && b.squares[row][column].Value == b.countAdjacentFlag(row, column) {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if i == 0 && j == 0 {
					continue
				}

				hasExploded := b.playRecursiveUtil(row+i, column+j, true)
				if hasExploded {
					b.revealAll()
					return PlayReturn{true, false, true}
				}
			}
		}
	}

	if b.numberDugTiles == b.numberRows*b.numberColumns-b.numberBombs {
		b.revealAll()
		return PlayReturn{true, true, false}
	}

	return PlayReturn{true, false, false}
}

// return true if bomb has exploded
func (b *Board) playRecursiveUtil(row, column int, canBombExplode bool) bool {
	if !b.isValid(row, column) {
		return false
	}

	if canBombExplode && !b.squares[row][column].isFlagged() && b.squares[row][column].IsBomb {
		return true
	}

	if b.squares[row][column].isFlagged() || b.squares[row][column].isRevealed() {
		return false
	}

	b.squares[row][column].reveal()
	b.numberDugTiles++

	if b.squares[row][column].Value != 0 {
		return false
	}

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			b.playRecursiveUtil(row+i, column+j, canBombExplode)
		}
	}

	return false
}

func (b *Board) revealAll() {
	for _, row := range b.squares {
		for _, square := range row {
			square.reveal()
		}
	}
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

			if b.squares[row+i][column+j].IsBomb {
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
