package util

func IsItemPresentInArray(item string, array []string) bool {
	for _, element := range array {
		if element == item {
			return true
		}
	}
	return false
}
