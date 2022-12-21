package game

const (
	Hidden   = 0
	Revealed = 1
	Flagged  = 2
)

type Square struct {
	state  int
	value  int
	isBomb bool
}

func newSquare(isBomb bool) *Square {
	return &Square{
		state:  Hidden,
		value:  0,
		isBomb: isBomb,
	}
}

func (s *Square) isRevealed() bool {
	return s.state == Revealed
}

func (s *Square) isFlagged() bool {
	return s.state == Flagged
}

func (s *Square) isHidden() bool {
	return s.state == Hidden
}

func (s *Square) incrementValue() {
	s.value++
}

func (s *Square) reveal() {
	s.state = Revealed
}

func (s *Square) flag() {
	if s.isRevealed() {
		return
	}

	s.state = Flagged
}

func (s *Square) unflag() {
	if s.isRevealed() {
		return
	}

	s.state = Hidden
}

func (s *Square) setValue(value int) {
	s.value = value
}

func (s *Square) bomb() {
	s.isBomb = true
}
