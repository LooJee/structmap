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

	mp, err := StructToMap(&st)
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

	m, err := StructToMap(st)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}

func TestDecode3(t *testing.T) {
	i := 1

	_, err := StructToMap(&i)
	if err != ErrNeedStruct {
		t.Fail()
	}
}

func TestDecode4(t *testing.T) {
	st := Bar{
		X: "hello",
		Y: "world",
		z: "zzz",
	}

	m, err := StructToMap(&st)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(m)
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

	dict, err := StructToMap(&st)
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

	_, err := StructToMap(sli)
	if err != ErrNeedStruct {
		t.Fatalf("want error : %v, got : %v", ErrNeedStruct, err)
	}
}

type complexType struct {
	A string
	B int64
	C float32
	D *int
	E *float32
	F *string
	G bool
	H *bool
	I []Foo
}

func TestComplextDecode(t *testing.T) {
	v := complexType{
		A: "hello",
		B: 11,
		C: 11.1,
		D: new(int),
		E: new(float32),
		F: new(string),
		G: true,
		I: []Foo{
			{X: "X", Y: "Y"},
			{X: "x", Y: "y"},
		},
	}

	*v.D = 1
	*v.E = 1.1
	*v.F = "world"

	dict, err := StructToMap(v)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(dict)
}

func TestMapToStruct(t *testing.T) {
	var foo complexType

	a := 1
	m := map[string]interface{}{
		"A": "123",
		"B": 234,
		"C": 1.1,
		"D": &a,
		"G": true,
		"I": []Foo{
			{X: "X", Y: "Y"},
			{X: "x", Y: "y"},
		},
	}

	if err := MapToStruct(m, foo); err != ErrNotPtr {
		t.Fatal("should be ErrNotPtr")
	}

	if err := MapToStruct(m, &foo); err != nil {
		t.Fatalf("should not be error, but got error : %v", err)
	}

	t.Logf("%#v", foo)
}
