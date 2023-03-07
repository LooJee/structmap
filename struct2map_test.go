package structmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	X string `stm:"x"`
	Y string `stm:"y"`
}

type Bar struct {
	X string
	Y string
	z string
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

	m, err := Decode(st)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
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
		z: "zzz",
	}

	m, err := Decode(&st)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m)
}

func TestDecode5(t *testing.T) {
	type Another struct {
		X string `stm:""`
		Y string
	}
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
	type IgnoreExample struct {
		X string `stm:"-"`
		Y string `stm:"y"`
	}

	st := IgnoreExample{
		X: "hello",
		Y: "world",
	}

	dict, err := Decode(&st)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := dict["x"]; ok {
		t.Fatal("x should be ignored")
	}

	if y, ok := dict["y"]; !ok {
		t.Fatal("y should be visible")
	} else {
		assert.Equal(t, "world", y)
	}
}

func TestDecode7(t *testing.T) {
	sli := []string{"123", "234", "345"}

	_, err := Decode(sli)
	if err != ErrNotValidElem {
		t.Fatalf("want error : %v, got : %v", ErrNotValidElem, err)
	}
}
