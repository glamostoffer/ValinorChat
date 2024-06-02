package errlist

import "errors"

var (
	ErrStopTimeout              = errors.New("stop timeout")
	ErrInvalidAffectedRowsCount = errors.New("invalid affected rows count")
)
