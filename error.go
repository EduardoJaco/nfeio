package nfeio

type ErrImpl interface {
	Error() string
}

type ErrMessage struct {
	Message string `json:"message"`
}

func (err ErrMessage) Error() string {
	return err.Message
}
