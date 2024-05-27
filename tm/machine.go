package tm

import (
	"fmt"
	"github.com/inazak/turing-machine/tm/rule"
)

const (
	BLANK = '\b'
)

type Machine struct {
	tape  *tape
	conf  map[string]rule.Rule
	state string
}

type tape struct {
	// left[0] [1] ... [-1] head [-1] ... [1] right[0]
	head  rune
	left  []rune //the end of slice is the right edge
	right []rune //the end of slice is the left edge
}

func New() *Machine {
	return &Machine{
		tape: &tape{head: BLANK},
		conf: make(map[string]rule.Rule),
	}
}

func (m *Machine) GetState() string {
	return m.state
}

func (m *Machine) SetState(state string) {
	m.state = state
}

func (m *Machine) ReadHead() rune {
	return m.tape.head
}

func (m *Machine) WriteHead(r rune) {
	m.tape.head = r
}

func (m *Machine) MoveLeft() {
	m.tape.right = append(m.tape.right, m.tape.head)
	if len(m.tape.left) > 0 {
		m.tape.head = m.tape.left[len(m.tape.left)-1]
		m.tape.left = m.tape.left[:len(m.tape.left)-1]
	} else {
		m.tape.head = BLANK
	}
}

func (m *Machine) MoveRight() {
	m.tape.left = append(m.tape.left, m.tape.head)
	if len(m.tape.right) > 0 {
		m.tape.head = m.tape.right[len(m.tape.right)-1]
		m.tape.right = m.tape.right[:len(m.tape.right)-1]
	} else {
		m.tape.head = BLANK
	}
}

func (m *Machine) HeadIndex() int {
	return len(m.tape.left)
}

func (m *Machine) Tape() []rune {
	lsize := len(m.tape.left)
	rsize := len(m.tape.right)
	runes := make([]rune, lsize+rsize+1)
	for i := 0; i < lsize; i++ {
		runes[i] = m.tape.left[i]
	}
	runes[lsize] = m.tape.head
	for i := 0; i < rsize; i++ {
		runes[lsize+1+i] = m.tape.right[rsize-i-1]
	}
	return runes
}

func (m *Machine) Rule() []rule.Rule {
	rules := []rule.Rule{}
	for _, r := range m.conf {
		rules = append(rules, rule.New(r.State, r.Match, r.Next, r.Write, r.Move))
	}
	return rules
}

func (m *Machine) AddRule(r rule.Rule) error {
	if _, ok := m.conf[r.Key()]; ok {
		return fmt.Errorf("rule's condition is conflicting")
	}
	m.conf[r.Key()] = r
	return nil
}

func (m *Machine) Step() bool {
	key := rule.CalculateKey(m.state, m.tape.head)
	r, ok := m.conf[key]
	if !ok {
		return false
	}
	m.state = r.Next
	m.tape.head = r.Write
	if r.IsMoveLeft() {
		m.MoveLeft()
	}
	if r.IsMoveRight() {
		m.MoveRight()
	}
	return true
}

func (m *Machine) Run() {
	for {
		if ok := m.Step(); !ok {
			break
		}
	}
}
