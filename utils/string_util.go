package utils

func IsEmpty(data string) bool {
	if len(data) < 1 {
		return true
	}

	return false
}
