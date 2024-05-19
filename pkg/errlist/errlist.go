package errlist

import "errors"

var (
	ErrInvalidAffectedRowsCount = errors.New("invalid affected rows count")
)
