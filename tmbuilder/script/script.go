package script

import (
	"fmt"
)

type Statement interface {
	StatementString() string
}

type RuleStatement struct {
	State  string
	Match  rune
	Write  rune
	Move   string
	Next   string
	Lineno int
}

type BeginStatement struct {
	State string
}

func (r RuleStatement) StatementString() string {
	return fmt.Sprintf("state %s match %s write %s move %s next %s",
		r.State, string(r.Match), string(r.Write), r.Move, r.Next)
}

func (b BeginStatement) StatementString() string {
	return fmt.Sprintf("start %s", b.State)
}
