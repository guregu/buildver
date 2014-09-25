package buildver

import (
	"testing"
)

// func TestClipZeroes(t *testing.T) {
// 	a, _ := New("1.2.3.0")
// 	b, _ := New("1.0.0.0")
// 	c, _ := New("1.2.0.1")

// 	t.Logf("%#v %#v %#v", a, b, c)
// 	t.Fail()
// }

func TestLess(t *testing.T) {
	a, _ := New("1.2.3.0")
	b, _ := New("1.0.0.0")

	if !b.Less(a) {
		t.Error("1.0.0.0 < 1.2.3.0")
	}
	if a.Less(b) {
		t.Error("1.0.0.0 < 1.2.3.0")
	}

	c, _ := New("1.0")
	d, _ := New("0.9")

	if !d.Less(c) {
		t.Error("0.9 < 1.0")
	}
	if c.Less(d) {
		t.Error("0.9 < 1.0")
	}
}

func TestVsInvalid(t *testing.T) {
	a, _ := New("dog")
	b, _ := New("4.1.10")

	if !a.Less(b) {
		t.Error("dog < 4.1.0")
	}
}

func TestMore(t *testing.T) {
	a, _ := New("4.1.10")
	b, _ := New("4.1.10.5")

	if !a.Less(b) {
		t.Error("4.1.10 < 4.1.10.5")
	}
}

func TestString(t *testing.T) {
	a, _ := New("dog")
	if str := a.String(); str != "0" {
		t.Error(str, "≠", "0")
	}

	b, _ := New("4.1.10")
	if str := b.String(); str != "4.1.10" {
		t.Error(str, "≠", "4.1.10")
	}
}
