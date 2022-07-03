package services

func AreFriends(strA, strB string) bool {
	if len(strA) != len(strB) {
		return false
	}
	for i := 0; i < len(strA); i++ {
		if strA == strB {
			return true
		}
		strB = strB[1:] + strB[:1]
	}
	return false
}
