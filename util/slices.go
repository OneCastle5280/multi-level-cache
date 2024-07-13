package util

// Contains
//
//	@Description: 判断 source 中是否存在元素 elem。
//	@param source
//	@param elem
//	@return bool
func Contains[E comparable](source []E, elem E) bool {
	for _, it := range source {
		if it == elem {
			return true
		}
	}
	return false
}
