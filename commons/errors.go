package commons

type errIterEnded struct{}

func (errIterEnded) Error() string {
	return "Iterator ended"
}

type errFuncIsNil struct{}

func (errFuncIsNil) Error() string {
	return "Function is nil"
}

var ErrIterEnded = errIterEnded{}
var ErrFuncIsNil = errFuncIsNil{}
