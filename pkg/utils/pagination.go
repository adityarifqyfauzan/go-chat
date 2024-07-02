package utils

// return limit, offset
func GetLimitOffset(page, size int) (int, int) {
	if page == 0 || size == 0 {
		return -1, -1
	}
	return size, (page - 1) * size
}
