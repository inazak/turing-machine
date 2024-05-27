package tm

import (
	"github.com/inazak/turing-machine/tm/rule"
	"testing"
)

func verify(m *Machine, want string) (string, bool) {
	runes := m.Tape()
	for i, r := range runes {
		if r == BLANK {
			runes[i] = '_'
		}
	}
	return string(runes), string(runes) == want
}

func TestBinaryIncrement(t *testing.T) {

	ma := New()
	ma.SetState("q0")
	ma.AddRule(rule.New("q0", BLANK, "q1", '1', "left"))
	ma.AddRule(rule.New("q1", BLANK, "q2", '1', "left"))
	ma.AddRule(rule.New("q2", BLANK, "q3", '0', "left"))
	ma.AddRule(rule.New("q3", BLANK, "q4", '1', "right"))
	ma.AddRule(rule.New("q4", '0', "q4", '0', "right"))
	ma.AddRule(rule.New("q4", '1', "q4", '1', "right"))
	ma.AddRule(rule.New("q4", BLANK, "q5", BLANK, "left"))

	ma.Run()

	if got, ok := verify(ma, "1011_"); !ok {
		t.Errorf("unexpected tape dump, got=%s", got)
	}
	if ma.HeadIndex() != 3 {
		t.Errorf("unexpected tape index, got=%d", ma.HeadIndex())
	}

	ma.SetState("p0")
	ma.AddRule(rule.New("p0", '0', "p1", '1', "right"))
	ma.AddRule(rule.New("p0", '1', "p0", '0', "left"))
	ma.AddRule(rule.New("p0", BLANK, "p1", '1', "right"))
	ma.AddRule(rule.New("p1", '0', "p1", '0', "right"))
	ma.AddRule(rule.New("p1", '1', "p1", '1', "right"))
	ma.AddRule(rule.New("p1", BLANK, "p2", BLANK, "left"))

	ma.Run()

	if got, ok := verify(ma, "1100_"); !ok {
		t.Errorf("unexpected tape dump, got=%s", got)
	}
	if ma.HeadIndex() != 3 {
		t.Errorf("unexpected tape index, got=%d", ma.HeadIndex())
	}
}
