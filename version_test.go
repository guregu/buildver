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

func TestLessWithEquals(t *testing.T) {
	a, _ := New("1.2")
	b, _ := New("1.2")
	if a.Less(b) {
		t.Error("1.2 = 1.2")
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

func TestEquals(t *testing.T) {
	a, _ := New("4.1.10.5")
	b, _ := New("4.1.10.5")

	if !a.Equals(b) || !b.Equals(a) {
		t.Error("4.1.10.5 = 4.1.10.5")
	}

	c, _ := New("1.0.0.0.0")
	d, _ := New("1")

	if !c.Equals(d) || !d.Equals(c) {
		t.Error("1.0.0.0.0 = 1")
	}

	e, _ := New("1.2")
	f, _ := New("1")

	if e.Equals(f) {
		t.Error("1.2 ≠ 1")
	}

	g, _ := New("2.5")
	h, _ := New("2.4")

	if g.Equals(h) {
		t.Error("2.5 ≠ 2.4")
	}
}

func TestFromInts(t *testing.T) {
	a, _ := New("1.2.3.0")
	b := FromInts(1, 2, 3, 0)
	if !a.Equals(b) {
		t.Error(`New("1.2.3.0") != FromInts(1, 2, 3, 0)`)
	}
}

// func TestContains(t *testing.T) {
// 	a := FromInts(1, 2, 3)
// 	b := FromInts(1, 2)
// 	c := FromInts(1)
// 	empty := FromInts()
// 	bad := FromInts(1, 9)

// 	if !a.Contains(b) {
// 		t.Error(a.String(), "should contain", b.String())
// 	}
// 	if !a.Contains(c) {
// 		t.Error(a.String(), "should contain", c.String())
// 	}

// 	if a.Contains(empty) {
// 		t.Error(a.String(), "should not contain an empty version")
// 	}

// 	if a.Contains(bad) {
// 		t.Error(a.String(), "should not contain", bad.String())
// 	}
// }

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

func TestExtraLetters(t *testing.T) {
	a, err := New("1.2a")
	if err != nil {
		t.Error("err should be nil but is", err)
	}
	b, _ := New("1.2")
	if a.Less(b) {
		t.Error("1.2a = 1.2")
	}
}
