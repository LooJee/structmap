package struct2map

import "errors"

var (
	ErrNotPtr       = errors.New("need a pointer")
	ErrNotValidElem = errors.New("pointer not point to struct")
	ErrNotValidTag  = errors.New("not valid tag")
	ErrNotValidKey  = errors.New("not valid key")
	ErrIgnore       = errors.New("ignore key")
	ErrNeedTag      = errors.New("need struct2map tag")
)
