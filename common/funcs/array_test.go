package funcs

import "testing"

func TestInArray(t *testing.T) {
	type args struct {
		item  string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: `test ""`, args: args{item: "", slice: []string{"", "1", "2"}}, want: true},
		{name: `test "0"`, args: args{item: "0", slice: []string{"", "1", "2"}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InArray(tt.args.item, tt.args.slice); got != tt.want {
				t.Errorf("In_array() = %v, want %v", got, tt.want)
			}
		})
	}
}
