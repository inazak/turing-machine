package tm

import (
	"github.com/inazak/turing-machine/tm/rule"
	"github.com/inazak/turing-machine/tm/tape"
	"testing"
)

var BLC = tape.BLANKCELL

func TestBinaryIncrement(t *testing.T) {
	ma := New()
	ma.AddRule(rule.New(0, BLC, 1, '1', "Left", "tape setup"))
	ma.AddRule(rule.New(1, BLC, 2, '1', "Left", "tape setup"))
	ma.AddRule(rule.New(2, BLC, 3, '0', "Left", "tape setup"))
	ma.AddRule(rule.New(3, BLC, 4, '1', "Right", "tape setup"))
	ma.AddRule(rule.New(4, '0', 4, '0', "Right", "tape setup"))
	ma.AddRule(rule.New(4, '1', 4, '1', "Right", "tape setup"))
	ma.AddRule(rule.New(4, BLC, 5, BLC, "Left", "tape setup"))

	ma.Run(-1)
	ma.ClearRuleAndState()

	if ma.Tape() != "1011_" {
		t.Errorf("unexpected tape dump, got=%s", ma.Tape())
	}
	if ma.HeadIndex() != 3 {
		t.Errorf("unexpected tape index, got=%d", ma.HeadIndex())
	}

	ma.AddRule(rule.New(0, '0', 1, '1', "Right", "binary increment"))
	ma.AddRule(rule.New(0, '1', 0, '0', "Left", "binary increment"))
	ma.AddRule(rule.New(0, BLC, 1, '1', "Right", "binary increment"))
	ma.AddRule(rule.New(1, '0', 1, '0', "Right", "binary increment"))
	ma.AddRule(rule.New(1, '1', 1, '1', "Right", "binary increment"))
	ma.AddRule(rule.New(1, BLC, 2, BLC, "Left", "binary increment"))

	ma.Run(-1)

	if ma.Tape() != "1100_" {
		t.Errorf("unexpected tape dump, got=%s", ma.Tape())
	}
	if ma.HeadIndex() != 3 {
		t.Errorf("unexpected tape index, got=%d", ma.HeadIndex())
	}
}
