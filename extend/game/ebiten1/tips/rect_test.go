package tips

import "testing"

func TestCheckReactCollision(t *testing.T) {
	type args struct {
		r1 Rect
		r2 Rect
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test-true", args: args{Rect{2, 2, 2, 2}, Rect{2, 2, 1, 1}}, want: true},
		{name: "test-false", args: args{Rect{2, 2, 2, 2}, Rect{2, 2, 4, 4}}, want: false},
		{name: "test-true", args: args{Rect{60, 58, 351.70296702817933, 369}, Rect{3, 15, 390.5, 369}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckReactCollision(tt.args.r1, tt.args.r2); got != tt.want {
				t.Errorf("CheckReactCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
