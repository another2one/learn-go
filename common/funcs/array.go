package funcs

//判断文件夹是否存在
func In_array(item string, slice []string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
