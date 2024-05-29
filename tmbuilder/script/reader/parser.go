package reader

import (
	"fmt"
	"github.com/inazak/turing-machine/tmbuilder/script"
)

type Parser struct {
	l         *Lexer
	currToken Token
	nextToken Token
	errors    []string
	lineno    int
	appdata   *appdata
}

type appdata struct {
	prevstate string
	blank     rune
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:       l,
		errors:  nil,
		appdata: &appdata{blank: '\b'}, // == tm.BLANK
	}
	p.readToken() // set p.nextToken
	p.readToken() // set p.nextToken and p.currToken
	p.lineno = 1
	return p
}

func (p *Parser) GetError() []string {
	return p.errors
}

func (p *Parser) readToken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.nextToken()

	if p.currToken.Type == UNKNOWN {
		p.errors = append(p.errors, fmt.Sprintf("Lexer got unknown token:"))
		p.errors = append(p.errors, fmt.Sprintf("  %s", p.l.GetError()))
	}
}

func (p *Parser) addErrorMessage(expect string) {
	s := p.l.lookbackText()
	p.errors = append(p.errors, fmt.Sprintf("line: %d", p.lineno))
	p.errors = append(p.errors, fmt.Sprintf("reading text '%s' <-", shorten(s, 10)))
	p.errors = append(p.errors, fmt.Sprintf("parser %s, but got [%s] %s", expect, p.currToken.Type, p.currToken.Literal))
}

func shorten(s string, i int) string {
	if len(s) <= i {
		return s
	} else {
		return "..." + s[len(s)-i:len(s)]
	}
}

func (p *Parser) isTokenType(t TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) isTokenTypeAndLiteral(t TokenType, l string) bool {
	return p.currToken.Type == t && p.currToken.Literal == l
}

type errorMessage struct {
	m string
}

func captureErrorMessageFromPanic(p *Parser) {
	if rec := recover(); rec != nil {
		if em, ok := rec.(errorMessage); ok {
			p.addErrorMessage(em.m)
		} else {
			panic(rec)
		}
	}
}

type anyLiteralType struct{}

var anyLiteral anyLiteralType

func (p *Parser) matchOrThrow(t TokenType, l any, errmsg string) string {
	var ok bool
	switch v := l.(type) {
	case anyLiteralType:
		ok = p.isTokenType(t)
	case string:
		ok = p.isTokenTypeAndLiteral(t, v)
	default:
		panic("unexpected type")
	}
	if ok {
		lt := p.currToken.Literal
		p.readToken()
		return lt
	} else {
		panic(errorMessage{m: errmsg})
	}
}

func (p *Parser) Parse() []script.Statement {
	list := []script.Statement{}

	for {
		for p.isTokenType(EOL) {
			p.lineno += 1
			p.readToken()
		}
		if p.isTokenType(EOT) {
			return list
		}

		list = append(list, p.parseLine())

		if p.GetError() != nil {
			break
		}
	}
	return list
}

func (p *Parser) parseLine() script.Statement {
	var stmt script.Statement

	if !p.isTokenType(KEYWORD) {
		p.addErrorMessage("expect instruction keyword")
		return stmt
	}

	switch p.currToken.Literal {
	case "state":
		stmt = p.parseRule()
	case "match":
		stmt = p.parseRule()
	case "begin":
		stmt = p.parseBegin()
	default:
		p.addErrorMessage("expect allowed instruction keyword")
	}

	if p.GetError() != nil {
		return stmt
	}
	if !p.isTokenType(EOT) && !p.isTokenType(EOL) {
		p.addErrorMessage("expect EOL or EOT")
		return stmt
	}
	return stmt
}

func firstRune(s string) rune {
	return []rune(s)[0]
}

func (p *Parser) parseRule() script.Statement {
	defer captureErrorMessageFromPanic(p)

	var state string
	var match rune
	var write rune
	var move string
	var next string

	if p.isTokenTypeAndLiteral(KEYWORD, "state") {
		p.readToken()
		state = p.matchOrThrow(SYMBOL, anyLiteral, "expect state symbol")
		p.appdata.prevstate = state
	} else {
		if p.appdata.prevstate == "" {
			p.addErrorMessage("expect keyword \"state\"")
			return nil
		}
		state = p.appdata.prevstate
	}

	_ = p.matchOrThrow(KEYWORD, "match", "expect keyword \"match\"")
	if p.isTokenTypeAndLiteral(KEYWORD, "blank") {
		p.readToken()
		match = p.appdata.blank
	} else {
		match = firstRune(p.matchOrThrow(CHAR, anyLiteral, "expect match charactor"))
	}

	_ = p.matchOrThrow(KEYWORD, "write", "expect keyword \"write\"")
	if p.isTokenTypeAndLiteral(KEYWORD, "blank") {
		p.readToken()
		write = p.appdata.blank
	} else {
		write = firstRune(p.matchOrThrow(CHAR, anyLiteral, "expect write charactor"))
	}

	_ = p.matchOrThrow(KEYWORD, "move", "expect keyword \"move\"")
	if p.isTokenTypeAndLiteral(KEYWORD, "left") {
		p.readToken()
		move = "left"
	} else if p.isTokenTypeAndLiteral(KEYWORD, "right") {
		p.readToken()
		move = "right"
	} else {
		p.addErrorMessage("expect direction \"left\" or \"right\"")
		return nil
	}

	_ = p.matchOrThrow(KEYWORD, "next", "expect keyword \"next\"")
	next = p.matchOrThrow(SYMBOL, anyLiteral, "expect next state symbol")

	return script.RuleStatement{
		State:  state,
		Match:  match,
		Write:  write,
		Move:   move,
		Next:   next,
		Lineno: p.lineno,
	}
}

func (p *Parser) parseBegin() script.Statement {
	defer captureErrorMessageFromPanic(p)

	var state string

	p.readToken()
	state = p.matchOrThrow(SYMBOL, anyLiteral, "expect state symbol")

	return script.BeginStatement{State: state}
}
