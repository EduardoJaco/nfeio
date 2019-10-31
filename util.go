package nfeio

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
