package main

type Token struct {
	code    int
	timeout int
	err     error
}

func (t Token) Code() int {
	return t.code
}

func (t Token) Timeout() int {
	return t.timeout
}

func (t Token) Error() error {
	return t.err
}

type Interface interface {
	Code() int
	Timeout() int
	Error() error
}
