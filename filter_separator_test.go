package logen

import "testing"

func Test_separatorFilter(t *testing.T) {
	type args struct {
		startMark string
		endMark   string
		s         string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "simple case",
			args: args{
				startMark: "[",
				endMark:   "]",
				s:         "hello [world] this is a test",
			},
			want: "hello  this is a test",
		},
		{
			name: "no end filter",
			args: args{
				startMark: "[",
				endMark:   "]",
				s:         "hello [world this is a test",
			},
			want: "hello [world this is a test",
		},
		{
			name: "multicharacter filter",
			args: args{
				startMark: "start",
				endMark:   "end",
				s:         "hello startworldend this is a test",
			},
			want: "hello  this is a test",
		},
		{
			name: "same start and end mark",
			args: args{
				startMark: "'",
				endMark:   "'",
				s:         "hello 'world' this is a test",
			},
			want: "hello  this is a test",
		},
		{
			name: "no filter in string",
			args: args{
				startMark: "[",
				endMark:   "]",
				s:         "hello world this is a test",
			},
			want: "hello world this is a test",
		},
		{
			name: "multiple filters",
			args: args{
				startMark: "[",
				endMark:   "]",
				s:         "hello [world] this [is] a [test]",
			},
			want: "hello  this  a ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separatorFilter(tt.args.startMark, tt.args.endMark)(tt.args.s); got != tt.want {
				t.Errorf("separatorFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
