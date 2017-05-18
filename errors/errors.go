package errors

type Error struct {
	Code   string `json:"code"`
	Entity string `json:"entity"`
}

type CustomError interface {
	Error() string
	ErrorCode() string
}

func NewError(code string, error error) *Error {
	return &Error{Code: code, Entity: error.Error()}
}

func (s *Error) ErrorCode() string {
	return s.Code
}

func (s *Error) Error() string {
	return s.Entity
}
