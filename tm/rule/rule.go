package rule

import (
	"fmt"
)

type Rule struct {
	conditionState int
	conditionCell  rune
	writeState     int
	writeCell      rune
	move           string
	comment        string
}

func New(cs int, cc rune, ws int, wc rune, mo string, co string) *Rule {
	r := &Rule{
		conditionState: cs,
		conditionCell:  cc,
		writeState:     ws,
		writeCell:      wc,
		move:           mo,
		comment:        co,
	}
	if !r.IsMoveLeft() && !r.IsMoveRight() {
		panic("unknown move direction")
	}
	return r
}

func (r *Rule) Key() string {
	return CalculateKey(r.conditionState, r.conditionCell)
}

func CalculateKey(state int, cell rune) string {
	return fmt.Sprintf("%v/%v", state, cell)
}

func (r *Rule) GetWriteState() int {
	return r.writeState
}

func (r *Rule) GetWriteCell() rune {
	return r.writeCell
}

func (r *Rule) IsMoveLeft() bool {
	switch r.move {
	case "left", "LEFT", "Left", "L":
		return true
	default:
		return false
	}
}

func (r *Rule) IsMoveRight() bool {
	switch r.move {
	case "right", "RIGHT", "Right", "R":
		return true
	default:
		return false
	}
}
