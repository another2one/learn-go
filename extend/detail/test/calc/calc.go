package calc

func Calc(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
