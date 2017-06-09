package future

import "fmt"

type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f
	return s
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

func (s *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		} else {
			s.successFunc(str)
		}
	}(s)
}

func setContext(msg string) ExecuteStringFunc {
	msg = fmt.Sprintf("%d Closure!\n", msg)

	return func() (string, error) {
		return msg, nil
	}
}

func init() {
	f := &MaybeString{}
	f.Success(func(s string) {
		fmt.Println("Success")
	}).Fail(func(e error) {
		fmt.Println("Fail")
	}).Execute(setContext("Hello"))
}
