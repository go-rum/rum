package errno

var (
	InternalServerError = &Errno{Code: 40001, Message: "internal server error"}
)

type Errno struct {
	Code    int
	Message string
}

func (err *Errno) Error() string {
	return err.Message
}

func New(code int, message string) *Errno {
	return &Errno{
		Code:    code,
		Message: message,
	}
}

func Decode(err error) (int, string) {
	if err == nil {
		return 0, ""
	}

	switch t := err.(type) {
	case *Errno:
		return t.Code, t.Message
	default:
		return InternalServerError.Code, InternalServerError.Message
	}
}
