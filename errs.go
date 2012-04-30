package expr

import "errors"

var (
	UnknownTokenErr = errors.New("don't know what to do with token")
	UnknownOpErr    = errors.New("don't know what to do with operator")
	UnknownLitErr   = errors.New("don't know what to do with literal")
)