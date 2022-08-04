package main

import "testing"

func Test_flippingBits(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Testing bit Flip",
			args: args{
				n: 1,
			},
			want: 4294967294,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flippingBits(tt.args.n); got != tt.want {
				t.Errorf("flippingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
