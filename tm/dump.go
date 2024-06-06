package tm

import (
	"fmt"
	"strings"
)

func (m *Machine) Dump(indent string) string {
	runes := m.Tape()
	for i, r := range runes {
		if r == BLANK {
			runes[i] = '_'
		}
	}

	dump := fmt.Sprintf("%s[%s]\n", indent, m.GetState())
	dump += fmt.Sprintf("%s%s\n", indent, string(runes))
	dump += fmt.Sprintf("%s%s^HEAD\n", indent, strings.Repeat(" ", m.HeadIndex()))
	return dump
}
