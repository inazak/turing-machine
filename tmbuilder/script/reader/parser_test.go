package reader

import (
	"github.com/inazak/turing-machine/tm"
	. "github.com/inazak/turing-machine/tmbuilder/script"
	"testing"
)

func helper(t *testing.T, text string, want []Statement) {
	l := NewLexer(text)
	p := NewParser(l)
	list := p.Parse()

	if errmsg := p.GetError(); errmsg != nil {
		t.Fatalf("%s", errmsg)
	}

	if len(want) != len(list) {
		t.Fatalf("statement count want=%d, but got=%d", len(want), len(list))
	}

	for i, _ := range list {
		w := want[i].StatementString()
		g := list[i].StatementString()
		if w != g {
			t.Fatalf("statement unmatch want=%s, but got=%s", w, g)
		}
	}
}

func TestParse(t *testing.T) {

	text := `
		#this is parser test1
		begin q0
		state q0 match 1 write 4 move left next q1
		state q1 match 2 write 5 move left  next q2
		         match 3 write 6 move right next q3
		state q3 match blank write blank move right next q4
	`

	want := []Statement{
		BeginStatement{State: "q0"},
		RuleStatement{State: "q0", Match: '1', Write: '4', Move: "left", Next: "q1"},
		RuleStatement{State: "q1", Match: '2', Write: '5', Move: "left", Next: "q2"},
		RuleStatement{State: "q1", Match: '3', Write: '6', Move: "right", Next: "q3"},
		RuleStatement{State: "q3", Match: tm.BLANK, Write: tm.BLANK, Move: "right", Next: "q4"},
	}

	helper(t, text, want)
}
