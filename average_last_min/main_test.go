package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		X int
		B []int
		Z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Testing solution",
			args: args{
				100, []int{10, 6, 6, 8}, 2,
			},
			want: 10,
		},
		{
			name: "Testing solution",
			args: args{
				100, []int{}, 0,
			},
			want: -1,
		},
		{
			name: "Testing solution",
			args: args{
				100, []int{6, 6, 8, 10, 12, 45}, 6,
			},
			want: 1,
		},
		{
			name: "Testing solution",
			args: args{
				100, []int{6, 6, 8, 10, 12, 1}, 6,
			},
			want: 8,
		},
		{
			name: "Testing solution",
			args: args{
				100, []int{6, 6, 8, 10, 12, 1}, 1,
			},
			want: 57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.X, tt.args.B, tt.args.Z); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
