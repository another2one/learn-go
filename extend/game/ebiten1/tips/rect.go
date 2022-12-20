package tips

type Rect struct {
	width, height, x, y float64
}

func CheckReactCollision(rMin, rMax Rect) bool {
	if rMin.width > rMax.width && rMin.height > rMax.width {
		// 这里有个问题 如果大的矩形移动速度超过了小的长或宽就可能检测不出来 （整个包在里面了）
		rMin, rMax = rMax, rMin
	}
	for _, point := range [4][2]float64{{rMin.x, rMin.y}, {rMin.x + rMin.width, rMin.y}, {rMin.x, rMin.y + rMin.height}, {rMin.x + rMin.width, rMin.y + rMin.height}} {
		if point[0] > rMax.x && point[0] < rMax.x+rMax.width && point[1] > rMax.y && point[1] < rMax.y+rMax.height {
			return true
		}
	}
	return false
}
