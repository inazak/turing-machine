package concept

import (
	"github.com/inazak/turing-machine/tm"
	"strings"
)

func verify(m *tm.Machine, want string) (string, bool) {
	runes := m.Tape()
	for i, r := range runes {
		if r == tm.BLANK {
			runes[i] = '_'
		}
	}
	s := strings.Trim(string(runes), "_")
	return s, s == want
}
