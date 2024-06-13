package concept

import (
	"github.com/inazak/turing-machine/tmbuilder"
	"testing"
)

func TestMatch(t *testing.T) {
	tape := "101101A1111101101B1111101101"
	want := "matched"

	// layout
	// -----
	// 101101A1111101101B1111101101
	// ^HEAD  |   |      |   |
	//        +---+      +---+
	//        count  ==  count
	//
	app := `
		state search-marka match 0     write 0     move right next search-marka
		                   match 1     write 1     move right next search-marka
		                   match A     write A     move right next check-a

		state check-a match 0     write 0     move right next move-b-final
		              match 1     write X     move right next move-markb
		              match X     write X     move right next check-a

		state move-markb match 0     write 0     move right next move-markb
		                 match 1     write 1     move right next move-markb
		                 match B     write B     move right next check-b

		state check-b match 0     write 0     move right next fail
		              match 1     write X     move right next next-marka
		              match X     write X     move right next check-b

		state next-marka match 0     write 0     move left  next next-marka
		                 match 1     write 1     move left  next next-marka
		                 match B     write B     move left  next next-marka
		                 match X     write X     move left  next next-marka
		                 match A     write A     move right next check-a

		state move-b-final match 0     write 0     move right next move-b-final
		                   match 1     write 1     move right next move-b-final
		                   match X     write X     move right next move-b-final
		                   match B     write B     move right next check-b-final

		state check-b-final match 0     write 0     move left  next matched
		                    match X     write X     move right next check-b-final
		                    match 1     write 1     move left  next fail
		                    match blank write blank move left  next fail

		#start point
		begin search-marka
	`

	b := tmbuilder.New()
	ma, err := b.Build(tape, app)
	if err != nil {
		t.Fatalf("builder.Build error, got=%s", err)
	}

	ma.Run()

	if ma.GetState() != want {
		t.Errorf("unexpected want=%s, but got=%s, tape=%s", want, ma.GetState(), tapedump(ma))
	}
}
