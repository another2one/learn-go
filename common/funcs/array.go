package funcs

// 判断 item 是否在 s1 里面
//  true 表示 item 在 s1 里面;  false 表示不在
//  暂时只支持字符串 后续使用泛型优化
func InArray(item string, s1 []string) bool {
	for _, v := range s1 {
		if v == item {
			return true
		}
	}
	return false
}
