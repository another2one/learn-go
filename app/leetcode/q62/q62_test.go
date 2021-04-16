package q62

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test1", args{3, 7}, 28},
		{"test2", args{3, 2}, 3},
		{"test3", args{7, 3}, 28},
		{"test4", args{3, 3}, 6},
		{"test5", args{3, 4}, 10},
		{"test6", args{4, 6}, 56},
		{"test7", args{9, 4}, 165},
		{"test8", args{11, 1}, 1},
		{"test9", args{11, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Count1 = 0
			Count2 = 0
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
			if got := uniquePaths1(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths1() = %v, want %v", got, tt.want)
			}
			if got := uniquePaths2(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths2() = %v, want %v", got, tt.want)
			}
			t.Logf("uniquePaths step:%d \n", Count1)
			t.Logf("uniquePaths2 step:%d \n", Count2)
		})
	}
}
