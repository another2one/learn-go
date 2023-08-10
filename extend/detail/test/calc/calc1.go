package calc

func Calc1(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}
