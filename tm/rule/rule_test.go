package rule

import (
	"testing"
)

func TestNew(t *testing.T) {
	r1 := New(1, 'a', 2, 'b', "left", "TestNew r1")
	r2 := New(1, 'b', 1, 'c', "right", "TestNew r2")

	if r1.Key() == r2.Key() {
		t.Errorf("keys are conflicting")
	}

	if r1.Key() != CalculateKey(1, 'a') {
		t.Errorf("keys do not match")
	}
}
