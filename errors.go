package structmap

import (
	"errors"
	"fmt"
)

var (
	ErrNotPtr     = errors.New("need a pointer")
	ErrNeedStruct = errors.New("pointer not point to struct")
)

type ErrTypeNotMatch struct {
	FieldName string
	WantType  string
	GotType   string
}

func (e *ErrTypeNotMatch) Error() string {
	return fmt.Sprintf("type of field : %s doesn't matched, want : %s, got %s", e.FieldName, e.WantType, e.GotType)
}
