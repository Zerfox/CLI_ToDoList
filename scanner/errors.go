package scanner

import "errors"

var ErrNeedExit = errors.New("exit")
var ErrEmptiInput = errors.New("empty input")
var ErrWrongArgs = errors.New("wrong arguments")
var ErrUnknownCommand = errors.New("unknown command")
var ErrScanError = errors.New("error scaner start")
