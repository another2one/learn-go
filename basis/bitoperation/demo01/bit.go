package demo01

// FIXME: 负数是补码操作

// 将某一位设置为1，例如设置第8位，从右向左数需要偏移7位,注意不要越界
// 1<<7=1000 0000 然后与a逻辑或|,偏移后的第8位为1，逻辑|运算时候只要1个为真就为真达到置1目的
//
// parameters:
//  p - 指定位置 0-7
func setAnyTo1(a int8, p int) int8 {
	p = p % 8

	var flag int8 = 1

	// 负数需要转化为正数操作
	if a < 0 {
		flag = -1
	}
	if p == 0 {
		if a > 0 {
			a = -a
		}
		return a
	}
	a = flag * a
	return flag * (a | 1<<(7-p))
}

// 将某一位设置为0，例如设置第4位，从右向左数需要偏移3位,注意不要越界
// 1<<3=0000 1000 然后取反得到 1111 0111 然后逻辑&a
//
// parameters:
//  p - 指定位置 0-7
func setAnyTo0(a int8, p int) int8 {
	p = p % 8

	var flag int8 = 1

	// 负数需要转化为正数操作
	if a < 0 {
		flag = -1
	}
	if p == 0 {
		if a < 0 {
			a = -a
		}
		return a
	}
	a = flag * a

	return flag * (a & (^(1 << (7 - p))))

}

// 获取某一位置的值: 即通过左右偏移来将将某位的值移动到第一位即可，当然也可以通过计算获得
// 如获取a的第4位先拿掉4位以上的值 a<<4=1110 0000,然后拿掉右边的3位即可 a>>7=0000 0001
//
// parameters:
//  p - 指定位置 0-7
func getAny(a int8, p int) int8 {
	p = p % 8
	if a < 0 {
		a = -a
	}
	if p == 0 {
		if a > 0 {
			return 0
		} else {
			return 1
		}
	}
	return (a << p) >> 7
}

// 取反某一位，即将某一位的1变0，0变1
//
// parameters:
//  p - 指定位置 0-7
func getAnyReverse(a int8, p int) int8 {
	p = p % 8
	return a ^ 1<<(7-p)
}

// 若tcp协议需要客户端先发送握手包，该包占用1个字节，其中前2位保留字段必须要为0，中间3位客户端对服务器版本要求，最后位客户端端版本
// 假设我们对服务器的版本要求和自己的版本都是3，那么我们该怎样构建这个包呢? 目标0001 1011
// 很多语言类型都没有直接 bit 单位，只要字节因此需要通过其他方法来得到,其实简单|或运算加上偏移即可,值得注意的网络使用的都是大端字节，传输前需要转换
func getTcpBit(remain, mid, last uint8) uint8 {
	return remain | (mid << 3) | (last)
}

// 求二进制数组的值 [8]uint8{0,0,0,0,0,1,1,0} = 6
func getBinaryValue(bs [8]uint8) int8 {
	var s int8
	var flag int8
	for i, v := range bs {
		v = v % 2
		// 第一位是符号
		if i == 0 {
			if v == 0 {
				flag = 1
			} else {
				flag = -1
			}
		} else {
			if v > 0 {
				s += 1 << (7 - i)
			}
		}

	}
	return s * flag
}
