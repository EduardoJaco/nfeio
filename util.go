package nfeio

import "strings"

func In(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return true
		}
	}

	return false

}

func NotIn(value interface{}, list ...interface{}) bool {

	for _, element := range list {
		if value == element {
			return false
		}
	}

	return true

}

func NotLikeAny(value string, list ...string) bool {

	for _, element := range list {
		if strings.Contains(value, element) {
			return false
		}
	}

	return true
}

func Ternary(condition bool, x, y interface{}) interface{} {

	if condition {
		return x
	}

	return y

}
