package concept

import (
	"github.com/inazak/turing-machine/tmbuilder"
	"testing"
)

func TestDelete(t *testing.T) {
	tape := "101101M1111101101"
	want := "101101M1101"

	// before
	// ------
	// 101101M1111101101
	// ^HEAD  |    |
	//        +----+
	//        delete
	//
	// after
	// -----
	// 101101M1101
	//       ^HEAD
	//
	app := `
		state search-mark-0 match 0     write 0     move right next search-mark-0
		                    match 1     write 1     move right next search-mark-0
		                    match M     write M     move right next move-rightest-0

		state move-rightest-0 match 0     write 0     move right next move-rightest-0
		                      match 1     write 1     move right next move-rightest-0
		                      match blank write blank move left  next del-start

		state del-start match 0     write blank move left next del-step-0
		                match 1     write blank move left next del-step-1

		state del-step-0 match 0     write 0     move left next del-step-0
		                 match 1     write 0     move left next del-step-1
		                 match M     write M     move left next finish

		state del-step-1 match 0     write 1     move left next del-step-0
		                 match 1     write 1     move left next del-step-1
		                 match M     write M     move right next move-rightest-0

		#start point
		begin search-mark-0
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
