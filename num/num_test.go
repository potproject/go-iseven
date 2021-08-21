package num

import (
	"reflect"
	"testing"
)

func Test_num(t *testing.T) {
	tests := []struct {
		name string
		arg  int64
		want []string
	}{
		{
			name: "0",
			arg:  0,
			want: []string{"zero"},
		},
		{
			name: "1",
			arg:  1,
			want: []string{"one"},
		},
		{
			name: "19",
			arg:  19,
			want: []string{"nineteen"},
		},
		{
			name: "20",
			arg:  20,
			want: []string{"twenty"},
		},
		{
			name: "99",
			arg:  99,
			want: []string{"ninety-nine"},
		},
		{
			name: "100",
			arg:  100,
			want: []string{"one", "hundred"},
		},
		{
			name: "101",
			arg:  101,
			want: []string{"one", "hundred", "one"},
		},
		{
			name: "1111",
			arg:  1111,
			want: []string{"one", "thousand", "one", "hundred", "eleven"},
		},
		{
			name: "1234567890",
			arg:  1234567890,
			want: []string{"one", "billion", "two", "hundred", "thirty-four", "million", "five", "hundred", "sixty-seven", "thousand", "eight", "hundred", "ninety"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("num() = %v, want %v", got, tt.want)
			}
		})
	}
}
