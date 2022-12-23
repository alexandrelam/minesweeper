package game

const (
	Hidden   = 0
	Revealed = 1
	Flagged  = 2
)

type Square struct {
	State  int  `json:"state"`
	Value  int  `json:"value"`
	IsBomb bool `json:"isBomb"`
}

func newSquare(isBomb bool) *Square {
	return &Square{
		State:  Hidden,
		Value:  0,
		IsBomb: isBomb,
	}
}

func (s *Square) isRevealed() bool {
	return s.State == Revealed
}

func (s *Square) isFlagged() bool {
	return s.State == Flagged
}

func (s *Square) isHidden() bool {
	return s.State == Hidden
}

func (s *Square) incrementValue() {
	s.Value++
}

func (s *Square) reveal() {
	s.State = Revealed
}

func (s *Square) flag() {
	if s.isRevealed() {
		return
	}

	s.State = Flagged
}

func (s *Square) unflag() {
	if s.isRevealed() {
		return
	}

	s.State = Hidden
}

func (s *Square) setValue(value int) {
	s.Value = value
}

func (s *Square) bomb() {
	s.IsBomb = true
}
