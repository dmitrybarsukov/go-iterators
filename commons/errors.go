package commons

type errIterEnded struct{}

func (errIterEnded) Error() string {
	return "Iterator ended"
}

type errFuncIsNil struct{}

func (errFuncIsNil) Error() string {
	return "Function is nil"
}

type errArgumentInvalid struct {
	Description string
}

func (e errArgumentInvalid) Error() string {
	return "argument invalid: " + e.Description
}

var ErrIterEnded = errIterEnded{}
var ErrFuncIsNil = errFuncIsNil{}
var ErrArgumentStepIsZero = errArgumentInvalid{Description: "step is zero, so iterator is infinite"}
var ErrArgumentStepHasWrongSign = errArgumentInvalid{Description: "iterator step has wrong sign"}
