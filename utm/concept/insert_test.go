package concept

import (
	"github.com/inazak/turing-machine/tmbuilder"
	"testing"
)

func TestInsert(t *testing.T) {
	tape := "101101A1111101101B01101"
	want := "101101A1111101101B1111101101"

	// before
	// ------
	// 101101A1111101101B01101
	// ^HEAD  |   |
	//        +---+
	//        copy
	//
	// after
	// -----
	// 101101A1111101101B1111101101
	//      ^HEAD        |   |
	//                   +---+
	//                   insert
	//
	app := `
		state search-marka-0 match 0     write 0     move right next search-marka-0
		                     match 1     write 1     move right next search-marka-0
		                     match A     write A     move right next check-rest-0

		state check-rest-0 match 0     write 0     move left  next restore-0
		                   match 1     write X     move right next move-markb-0
		                   match X     write X     move right next check-rest-0

		state move-markb-0 match 0     write 0     move right next move-markb-0
		                   match 1     write 1     move right next move-markb-0
		                   match B     write B     move right next start-insert

		state start-insert match 0     write 1     move right next ins-step-0
		                   match 1     write 1     move right next ins-step-1

		state ins-step-0 match 0     write 0     move right next ins-step-0
		                 match 1     write 0     move right next ins-step-1
		                 match blank write 0     move left  next back-marka-0

		state ins-step-1 match 0     write 1     move right next ins-step-0
		                 match 1     write 1     move right next ins-step-1
		                 match blank write 1     move left  next back-marka-0

		state back-marka-0 match 0     write 0     move left  next back-marka-0
		                   match 1     write 1     move left  next back-marka-0
		                   match B     write B     move left  next back-marka-0
		                   match X     write X     move left  next back-marka-0
		                   match A     write A     move right next check-rest-0

		state restore-0 match X     write 1     move left next restore-0
		                match A     write A     move left next finish

		#start point
		begin search-marka-0
	`

	b := tmbuilder.New()
	ma, err := b.Build(tape, app)
	if err != nil {
		t.Fatalf("builder.Build error, got=%s", err)
	}

	ma.Run()

	if got, ok := verify(ma, want); !ok {
		t.Errorf("unexpected tape dump, want=%s, but got=%s", want, got)
	}
}
