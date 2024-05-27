package rule

import (
	"testing"
)

func TestNewAndKey(t *testing.T) {
	r1 := New("q1", 'a', "q2", 'b', "left")
	r2 := New("q1", 'b', "q1", 'c', "right")

	if r1.Key() == r2.Key() {
		t.Errorf("keys are conflicting")
	}

	if r1.Key() != CalculateKey("q1", 'a') {
		t.Errorf("keys do not match")
	}
}
