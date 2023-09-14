package errors

import "fmt"

type InvalidParam struct {
	Param string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("param '%s' is invalid", i.Param)
}

type Err struct {
	Code int
	Msg  string
}

func (i Err) Error() string {
	return i.Msg
}
