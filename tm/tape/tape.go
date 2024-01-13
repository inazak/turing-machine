package tape

const (
	BLANKCELL = '\b'
)

type Tape struct {
	// left[0] left[1] ... left[-1] head right[-1] ... right[1] right[0]
	head       rune
	left       []rune //the end of slice is the right edge
	right      []rune //the end of slice is the left edge
	BlankPrint rune
}

func Make(left []rune, head rune, right []rune) *Tape {
	for i, j := 0, len(right)-1; i < len(right)/2; i, j = i+1, j-1 {
		right[i], right[j] = right[j], right[i]
	}
	return &Tape{
		head:       head,
		left:       left,
		right:      right,
		BlankPrint: '_',
	}
}

func New() *Tape {
	return Make(nil, BLANKCELL, nil)
}

func (t *Tape) Read() rune {
	return t.head
}

func (t *Tape) Write(r rune) {
	t.head = r
}

func (t *Tape) MoveLeft() {
	t.right = append(t.right, t.head)
	if len(t.left) > 0 {
		t.head = t.left[len(t.left)-1]
		t.left = t.left[:len(t.left)-1]
	} else {
		t.head = BLANKCELL
	}
}

func (t *Tape) MoveRight() {
	t.left = append(t.left, t.head)
	if len(t.right) > 0 {
		t.head = t.right[len(t.right)-1]
		t.right = t.right[:len(t.right)-1]
	} else {
		t.head = BLANKCELL
	}
}

func (t *Tape) HeadIndex() int {
	return len(t.left)
}

func (t *Tape) dumpLeft() string {
	size := len(t.left)
	dump := make([]rune, size)
	for i := 0; i < size; i++ {
		if t.left[i] != BLANKCELL {
			dump[i] = t.left[i]
		} else {
			dump[i] = t.BlankPrint
		}
	}
	return string(dump)
}

func (t *Tape) dumpRight() string {
	size := len(t.right)
	dump := make([]rune, size)
	for i := 0; i < size; i++ {
		if t.right[size-i-1] != BLANKCELL {
			dump[i] = t.right[size-i-1]
		} else {
			dump[i] = t.BlankPrint
		}
	}
	return string(dump)
}

func (t *Tape) Dump() string {
	var c rune
	if t.head == BLANKCELL {
		c = t.BlankPrint
	} else {
		c = t.head
	}
	return t.dumpLeft() + string(c) + t.dumpRight()
}
