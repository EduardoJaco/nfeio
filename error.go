package nfeio

import "strings"

type ErrMessage struct {
	Message string `json:"message"`
}

type ErrArray struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrMessage) Error() string {
	return e.Message
}

func (e ErrArray) Error() string {

	var messages []string

	for _, err := range e.Errors {
		messages = append(messages, err.Message)
	}

	return strings.Join(messages, ",")

}

func (e ErrArray) Count() int {
	return len(e.Errors)
}
