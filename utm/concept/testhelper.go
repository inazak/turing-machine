package concept

import (
	"github.com/inazak/turing-machine/tm"
	"strings"
)

func tapedump(m *tm.Machine) string {
	runes := m.Tape()
	for i, r := range runes {
		if r == tm.BLANK {
			runes[i] = '_'
		}
	}
	return strings.Trim(string(runes), "_")
}

func verify(m *tm.Machine, want string) (string, bool) {
	s := tapedump(m)
	return s, s == want
}
