package concept

import (
	"github.com/inazak/turing-machine/tmbuilder"
	"testing"
)

func TestReplace(t *testing.T) {
	tape := "0101100011011100011110100010101000"
	want := "0101100011000111101000111010101000"

	// before
	// ------
	//            HEAD           RIGHT
	// 0101100011011100011110100010101000
	//         STATE    LEFT
	//
	// working
	// -----
	// 0101100011011100011110100010101000
	// ^HEAD      |  |           ^
	//            +--+           |
	//            copy,del       insert
	//
	// after
	// -----
	// 0101100011000111101000111010101000
	//
	app := `
		state s1-mark-a-0 match 0     write 0     move right next s1-mark-a-1
		                  match 1     write 1     move right next s1-mark-a-0

		state s1-mark-a-1 match 0     write 0     move right next s1-mark-a-2
		                  match 1     write 1     move right next s1-mark-a-0

		state s1-mark-a-2 match 0     write 0     move right next s1-mark-a-3
		                  match 1     write 1     move right next s1-mark-a-0

		state s1-mark-a-3 match 0     write A     move right next s1-mark-b-0
		                  match 1     write 1     move right next s1-mark-a-3


		state s1-mark-b-0 match 0     write 0     move right next s1-mark-b-1
		                  match 1     write 1     move right next s1-mark-b-0

		state s1-mark-b-1 match 0     write 0     move right next s1-mark-b-2
		                  match 1     write 1     move right next s1-mark-b-0

		state s1-mark-b-2 match 0     write 0     move right next s1-mark-b-3
		                  match 1     write 1     move right next s1-mark-a-0

		state s1-mark-b-3 match 0     write 0     move right next s1-mark-b-4
		                  match 1     write 1     move right next s1-mark-b-3

		state s1-mark-b-4 match 0     write 0     move right next s1-mark-b-5
		                  match 1     write 1     move right next s1-mark-b-3

		state s1-mark-b-5 match 0     write B     move right next s1-insert-zero
		                  match 1     write 1     move right next s1-mark-b-3

		state s1-insert-zero match 0     write 0     move right next s1-insert-zero-0
		                     match 1     write 0     move right next s1-insert-zero-1

		state s1-insert-zero-0 match 0     write 0     move right next s1-insert-zero-0
		                       match 1     write 0     move right next s1-insert-zero-1
		                       match blank write 0     move left  next s1-back-mark-a

		state s1-insert-zero-1 match 0     write 1     move right next s1-insert-zero-0
		                       match 1     write 1     move right next s1-insert-zero-1
		                       match blank write 1     move left  next s1-back-mark-a

		state s1-back-mark-a match 0     write 0     move left  next s1-back-mark-a
		                     match 1     write 1     move left  next s1-back-mark-a
		                     match X     write X     move left  next s1-back-mark-a
		                     match B     write B     move left  next s1-back-mark-a
		                     match A     write A     move right next s1-do-or-not

		state s1-move-to-b match 0     write 0     move right next s1-move-to-b
		                   match 1     write 1     move right next s1-move-to-b
		                   match B     write B     move right next s1-start-insert

		state s1-do-or-not match 0     write 0     move left  next s1-restore-0
		                   match 1     write X     move right next s1-move-to-b
		                   match X     write X     move right next s1-do-or-not

		state s1-start-insert match 0     write 1     move right next s1-insert-has-0
		                      match 1     write 1     move right next s1-insert-has-1

		state s1-insert-has-0 match 0     write 0     move right next s1-insert-has-0
		                      match 1     write 0     move right next s1-insert-has-1
		                      match blank write 0     move left  next s1-back-mark-a

		state s1-insert-has-1 match 0     write 1     move right next s1-insert-has-0
		                      match 1     write 1     move right next s1-insert-has-1
		                      match blank write 1     move left  next s1-back-mark-a

		state s1-restore-0 match 0     write 0     move right next s1-restore-0
		                   match 1     write 1     move right next s1-restore-0
		                   match A     write A     move right next s1-restore-0
		                   match B     write B     move right next s1-restore-0
		                   match X     write X     move right next s1-restore-0
		                   match blank write blank move left  next s1-restore-1

		state s1-restore-1 match 0     write 0     move left  next s1-restore-1
		                   match 1     write 1     move left  next s1-restore-1
		                   match X     write 1     move left  next s1-restore-1
		                   match B     write 0     move left  next s1-restore-1
		                   match A     write A     move right next s2-rightest

		state s2-rightest match 0     write 0     move right next s2-rightest
		                  match 1     write 1     move right next s2-rightest
		                  match blank write blank move left  next s2-del-start

		state s2-del-start match 0     write blank move left next s2-del-step-0
		                   match 1     write blank move left next s2-del-step-1

		state s2-del-step-0 match 0     write 0     move left next s2-del-step-0
		                    match 1     write 0     move left next s2-del-step-1
		                    match A     write 0     move left next finish

		state s2-del-step-1 match 0     write 1     move left next s2-del-step-0
		                    match 1     write 1     move left next s2-del-step-1
		                    match A     write A     move right next s2-rightest

		#start point
		begin s1-mark-a-0
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
