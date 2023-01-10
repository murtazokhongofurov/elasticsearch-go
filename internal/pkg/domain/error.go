package domain

import (
	"errors"
)

var (
	errorNotFount = errors.New("not found")
	errorConflict = errors.New("conflict")
)
