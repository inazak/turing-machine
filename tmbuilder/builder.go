package tmbuilder

import (
	"fmt"
	"github.com/inazak/turing-machine/tm"
	"github.com/inazak/turing-machine/tm/rule"
	"github.com/inazak/turing-machine/tmbuilder/script"
	"github.com/inazak/turing-machine/tmbuilder/script/reader"
)

type Builder struct {
	tapePrefix string
}

func New() *Builder {
	return &Builder{
		tapePrefix: "tape-setup",
	}
}

func (b *Builder) SetTapePrefix(prefix string) {
	b.tapePrefix = prefix
}

// the position of head on the tape generated by Builder() is as follows.
// when there is no tape definition, the position of head is above the blank cell.
//   _
//   ^HEAD
//
// when tape is defined, the position of head is the leftmost cell containing the value.
//   _ 1 2 3
//     ^HEAD
//
func (b *Builder) Build(tapeDefinition, appDefinition string) (*tm.Machine, error) {

	// Fisrt, read the application definition in the parser
	// and add it as a rule to Machine
	l := reader.NewLexer(appDefinition)
	p := reader.NewParser(l)
	stmts := p.Parse()
	if errmsg := p.GetError(); errmsg != nil {
		return nil, fmt.Errorf("%v", errmsg)
	}

	ma := tm.New()
	for _, stmt := range stmts {
		switch v := stmt.(type) {
		case script.RuleStatement:
			err := ma.AddRule(rule.New(v.State, v.Match, v.Next, v.Write, v.Move))
			if err != nil {
				return nil, err
			}
		case script.BeginStatement:
			ma.SetState(v.State)
		default:
			return nil, fmt.Errorf("unknown script statement")
		}
	}
	if ma.GetState() == "" {
		return nil, fmt.Errorf("no define to begin state")
	}

	// Next, covert the tape definition to a rule
	// add it to Machine in same way
	runes := toReverseRunes(tapeDefinition)
	for i, r := range runes {
		curr := fmt.Sprintf("%s-%05d", b.tapePrefix, i)
		next := fmt.Sprintf("%s-%05d", b.tapePrefix, i+1)
		if err := ma.AddRule(rule.New(curr, tm.BLANK, next, r, "left")); err != nil {
			return nil, err
		}
	}

	// Finally, connect the end of tape creation rule to the beginning
	// of the application rule
	if tapeDefinition != "" {
		curr := fmt.Sprintf("%s-%05d", b.tapePrefix, len(runes))
		next := ma.GetState()
		if err := ma.AddRule(rule.New(curr, tm.BLANK, next, tm.BLANK, "right")); err != nil {
			return nil, err
		}
		ma.SetState(fmt.Sprintf("%s-%05d", b.tapePrefix, 0))
	}

	return ma, nil
}

func toReverseRunes(s string) []rune {
	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}
	size := len(runes)
	for i := 0; i < size/2; i += 1 {
		runes[i], runes[size-i-1] = runes[size-i-1], runes[i]
	}
	return runes
}
