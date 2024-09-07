package util

func GetPagination(page int, limit int) (int, int) {
	offset := 0
	if page != 0 && page != 1 {
		offset = (page - 1) * limit
	}

	return offset, limit
}