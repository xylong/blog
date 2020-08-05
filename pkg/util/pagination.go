package util

// GetPage 分页
func GetPage(page, size int) (offset uint) {
	if page > 0 {
		offset = uint((page - 1) * size)
	}
	return
}
