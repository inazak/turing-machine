package tm

import (
	"fmt"
	"github.com/inazak/turing-machine/tm/rule"
	"github.com/inazak/turing-machine/tm/tape"
)

type Machine struct {
	tape  *tape.Tape
	conf  map[string]*rule.Rule
	state int
}

func New() *Machine {
	return &Machine{
		tape: tape.New(),
		conf: make(map[string]*rule.Rule),
	}
}

func (m *Machine) AddRule(r *rule.Rule) error {
	if _, ok := m.conf[r.Key()]; ok {
		return fmt.Errorf("rule's condition is conflicting")
	}
	m.conf[r.Key()] = r
	return nil
}

func (m *Machine) ClearRuleAndState() {
	m.conf = make(map[string]*rule.Rule)
	m.state = 0
}

func (m *Machine) Step() bool {
	key := rule.CalculateKey(m.state, m.tape.Read())
	r, ok := m.conf[key]
	if !ok {
		return false
	}
	m.state = r.GetWriteState()
	m.tape.Write(r.GetWriteCell())
	if r.IsMoveLeft() {
		m.tape.MoveLeft()
	}
	if r.IsMoveRight() {
		m.tape.MoveRight()
	}
	return true
}

func (m *Machine) Run(breakWhenState int) {
	for {
		if ok := m.Step(); !ok {
			break
		}
		if m.state == breakWhenState {
			break
		}
	}
}

func (m *Machine) State() int {
	return m.state
}

func (m *Machine) HeadIndex() int {
	return m.tape.HeadIndex()
}

func (m *Machine) Tape() string {
	return m.tape.Dump()
}
