package demo02

import (
	"math/rand"
	"reflect"
	"testing"
)

type args struct {
	s1 []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"test1：", args{[]int{6, 7, 9, 10, 6}}, []int{6, 7, 9, 10}},
	{"test2：", args{[]int{6, 6, 9, 9, 6}}, []int{6, 9}},
}

func TestSliceToSet1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToSet1(tt.args.s1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToSet1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToSet2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToSet2(tt.args.s1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToSet2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToSet3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToSet3(tt.args.s1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToSet3() = %v, want %v", got, tt.want)
			}
		})
	}
}

var s1 = getS1()

func getS1() []int {
	start, end := 0, 1000
	s := make([]int, 1000)
	for start < end {
		s[start] = rand.Intn(end)
		start++
	}
	return s
}

func BenchmarkSliceToSet1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceToSet1(s1)
	}
}

func BenchmarkSliceToSet2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceToSet2(s1)
	}
}

func BenchmarkSliceToSet3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceToSet3(s1)
	}
}
