package errors

type errIterEnded struct{}

func (errIterEnded) Error() string {
	return "Iterator ended"
}

var ErrIterEnded = errIterEnded{}
