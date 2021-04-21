package demo01

import (
	"testing"
)

func Test_setAnyTo1(t *testing.T) {
	type args struct {
		a int8
		p int
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{"test0", args{getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 0, 1}), 0}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1})},
		{"test1", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 4}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1})},
		{"test2", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 1, 1})},
		{"test3", args{getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 1, 1})},
		{"test4", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 0}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setAnyTo1(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("setAnyTo1() = %08b, want %08b", got, tt.want)
			}
		})
	}
}

func Test_setAnyTo0(t *testing.T) {
	type args struct {
		a int8
		p int
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{"test1", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 4}, getBinaryValue([8]uint8{1, 0, 1, 1, 0, 1, 0, 1})},
		{"test2", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1})},
		{"test3", args{getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setAnyTo0(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("setAnyTo0() = %08b, want %08b", got, tt.want)
			}
		})
	}
}

func Test_getAny(t *testing.T) {
	type args struct {
		a int8
		p int
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{"test1", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 4}, 1},
		{"test2", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 6}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAny(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("getAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAnyReverse(t *testing.T) {
	type args struct {
		a int8
		p int
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{"test1", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 4}, getBinaryValue([8]uint8{1, 0, 1, 1, 0, 1, 0, 1})},
		{"test2", args{getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{1, 0, 1, 1, 1, 1, 1, 1})},
		{"test3", args{getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 0, 1}), 6}, getBinaryValue([8]uint8{0, 0, 1, 1, 1, 1, 1, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAnyReverse(tt.args.a, tt.args.p); got != tt.want {
				t.Errorf("getAnyReverse() = %08b, want %08b", got, tt.want)
			}
		})
	}
}

func Test_getTcpBit(t *testing.T) {
	type args struct {
		remain uint8
		mid    uint8
		last   uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
		{"test01", args{0, 3, 3}, uint8(getBinaryValue([8]uint8{0, 0, 0, 1, 1, 0, 1, 1}))},
		{"test01", args{0, 4, 1}, uint8(getBinaryValue([8]uint8{0, 0, 1, 0, 0, 0, 0, 1}))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTcpBit(tt.args.remain, tt.args.mid, tt.args.last); got != tt.want {
				t.Errorf("getTcpBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBinaryValue(t *testing.T) {
	type args struct {
		bs [8]uint8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		// TODO: Add test cases.
		{"test normal", args{[8]uint8{0, 0, 0, 0, 0, 1, 1, 0}}, 6},
		{"test negative", args{[8]uint8{1, 0, 0, 0, 0, 1, 1, 0}}, -6},
		{"test 0", args{[8]uint8{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBinaryValue(tt.args.bs); got != tt.want {
				t.Errorf("getBinaryValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
