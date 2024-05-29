package tmbuilder

import (
	"github.com/inazak/turing-machine/tm"
	"testing"
)

func verify(m *tm.Machine, want string) (string, bool) {
	runes := m.Tape()
	for i, r := range runes {
		if r == tm.BLANK {
			runes[i] = '_'
		}
	}
	return string(runes), string(runes) == want
}

func TestBinaryIncrement(t *testing.T) {
	tape := "1011"
	app := `
		#this is binary increment
		state moveright-q0 match 0     write 0     move right next moveright-q0
		state moveright-q0 match 1     write 1     move right next moveright-q0
		state moveright-q0 match blank write blank move left  next increment-q0

		state increment-q0 match 0     write 1     move right next increment-q1
		                   match 1     write 0     move left  next increment-q0
		                   match blank write 1     move right next increment-q1
		state increment-q1 match 0     write 0     move right next increment-q1
		                   match 1     write 1     move right next increment-q1
		                   match blank write blank move left  next increment-qf
		#start point
		begin moveright-q0
	`

	b := New()
	ma, err := b.Build(tape, app)
	if err != nil {
		t.Fatalf("builder.Build error, got=%s", err)
	}

	ma.Run()

	if got, ok := verify(ma, "_1100_"); !ok {
		t.Errorf("unexpected tape dump, got=%s", got)
	}
	if ma.HeadIndex() != 4 {
		t.Errorf("unexpected tape index, got=%d", ma.HeadIndex())
	}
}
