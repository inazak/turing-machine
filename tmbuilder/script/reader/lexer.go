package reader

func getKeyword() []string {
	return []string{
		"state",
		"match",
		"write",
		"move",
		"next",
		"right",
		"left",
		"blank",
		"begin",
	}
}

type Lexer struct {
	text         []rune
	line         int
	currPosition int
	nextPosition int
	c            rune //charactor of currPosition
	msg          string
}

func NewLexer(text string) *Lexer {
	l := &Lexer{text: []rune(text), line: 1}
	l.read()
	return l
}

func (l *Lexer) GetError() string {
	return l.msg
}

func (l *Lexer) lookbackText() string {
	if l.currPosition < 2 {
		return ""
	} else {
		return string(l.text[:l.currPosition-2])
	}
}

func (l *Lexer) read() {
	if l.nextPosition < len(l.text) {
		l.c = l.text[l.nextPosition]
		l.currPosition = l.nextPosition
		l.nextPosition += 1
	} else {
		l.c = 0
	}
}

func (l *Lexer) skipSpace() {
	for l.c == ' ' || l.c == '\t' {
		l.read()
	}
}

func (l *Lexer) skipComment() {
	if l.c == '#' {
		for {
			l.read()
			if l.c == '\r' || l.c == '\n' || l.c == 0 {
				break
			}
		}
	}
}

func (l *Lexer) readKeyword() (bool, string) {
	p := l.currPosition
	for p < len(l.text) && 'a' <= l.text[p] && l.text[p] <= 'z' {
		p += 1
	}
	if p < len(l.text) && (l.text[p] == '-' || l.text[p] == '_') {
		return false, ""
	}
	w := string(l.text[l.currPosition:p])
	for _, k := range getKeyword() {
		if w == k {
			l.currPosition = p - 1
			l.nextPosition = p
			l.c = l.text[l.currPosition]
			return true, k
		}
	}
	return false, ""
}

func (l *Lexer) nextIs(r rune) bool {
	return l.nextPosition < len(l.text) && l.text[l.nextPosition] == r
}

func (l *Lexer) isSymbolHead() bool {
	return ('a' <= l.c && l.c <= 'z')
}

func (l *Lexer) nextIsSymbolRest() bool {
	if l.nextPosition < len(l.text) {
		c := l.text[l.nextPosition]
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') || c == '-' || c == '_' {
			return true
		}
	}
	return false
}

func (l *Lexer) isSymbol() (bool, string) {
	if !l.isSymbolHead() {
		return false, ""
	}
	if !l.nextIsSymbolRest() {
		return false, ""
	}
	s := []rune{l.c}
	for l.nextIsSymbolRest() {
		l.read()
		s = append(s, l.c)
	}
	return true, string(s)
}

func (l *Lexer) isChar() (bool, string) {
	c := l.c
	if ' ' <= c && c <= '~' {
		return true, string(l.c)
	}
	return false, ""
}

func (l *Lexer) nextToken() Token {
	var tk Token
	l.skipSpace()
	l.skipComment()

	switch l.c {
	case 0:
		tk = NewToken(EOT, "")
	case '\r':
		if l.nextIs('\n') {
			l.read()
		}
		tk = NewToken(EOL, "")
	case '\n':
		tk = NewToken(EOL, "")
	default:
		if ok, s := l.readKeyword(); ok {
			tk = NewToken(KEYWORD, s)
		} else if ok, s := l.isSymbol(); ok {
			tk = NewToken(SYMBOL, s)
		} else if ok, s := l.isChar(); ok {
			tk = NewToken(CHAR, s)
		} else {
			tk = NewToken(UNKNOWN, string(l.c))
			l.msg = "unallowed character -> " + string(l.c)
		}
	}

	l.read()
	return tk
}
