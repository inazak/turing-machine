package tape

import (
	"testing"
)

func TestNewAndWrite(t *testing.T) {
	tape := New()

	if !tape.IsBlank() {
		t.Errorf("initial cell is not blank")
	}

	tape.Write('b')
	if tape.Read() != 'b' {
		t.Errorf("cannot write or read")
	}

	tape.MoveLeft()
	tape.Write('a')
	tape.MoveRight()
	tape.MoveRight()
	tape.Write('c')
	tape.MoveRight()
	tape.BlankPrint = '_'

	if tape.Dump() != "abc_" {
		t.Errorf("unexpected tape dump")
	}

	if tape.HeadIndex() != 3 {
		t.Errorf("unexpected index")
	}
}
