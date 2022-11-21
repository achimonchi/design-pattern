package future

type SuccessFunc func(string)
type FailFunc func(string)
type ExecuteStringFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (m *MaybeString) Success(f SuccessFunc) *MaybeString {
	m.successFunc = f
	return m
}

func (m *MaybeString) Fail(f FailFunc) *MaybeString {
	m.failFunc = f
	return m
}

func (m *MaybeString) Execute(f ExecuteStringFunc) {
	go func(s *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err.Error())
		} else {
			s.successFunc(str)
		}
	}(m)
}
