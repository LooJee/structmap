package structmap

import "testing"

type Foo struct {
	X string `struct2map:"key:x"`
	Y string `struct2map:"key:y"`
}

type Bar struct {
	X string
	Y string
}

type Another struct {
	X string `struct2map:""`
	Y string
}

type IgnoreExample struct {
	X string `struct2map:"-"`
	Y string `struct2map:"key:y"`
}

func TestDecode(t *testing.T) {
	st := Foo{
		X: "hello",
		Y: "world",
	}

	mp, err := Decode(&st)
	if err != nil {
		t.Fatal(err)
	}

	if mp["x"] == "hello" && mp["y"] == "world" {
		t.Log(mp)
	} else {
		t.Fatal(mp)
	}
}

func TestDecode2(t *testing.T) {
	st := Foo{
		X: "hello",
		Y: "world",
	}

	_, err := Decode(st)
	if err != ErrNotPtr {
		t.Fail()
	}
}

func TestDecode3(t *testing.T) {
	i := 1

	_, err := Decode(&i)
	if err != ErrNotValidElem {
		t.Fail()
	}
}

func TestDecode4(t *testing.T) {
	st := Bar{
		X: "hello",
		Y: "world",
	}

	_, err := Decode(&st)
	if err != ErrNeedTag {
		t.Fail()
	}
}

func TestDecode5(t *testing.T) {
	st := Another{
		X: "hello",
		Y: "world",
	}

	_, err := Decode(&st)
	if err != ErrNotValidTag {
		t.Fail()
	}
}

func TestDecode6(t *testing.T) {
	st := IgnoreExample{
		X: "hello",
		Y: "world",
	}

	dict, err := Decode(&st)
	if err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log(dict)
	}
}
