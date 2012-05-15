package riakpbc

import (
	"errors"
)

var (
	ErrLengthZero     = errors.New("length response 0")
	ErrCorruptHeader  = errors.New("corrupt header")
	ErrObjectNotFound = errors.New("object not found")
	ErrNoSuchCommand  = errors.New("no such command")
)
