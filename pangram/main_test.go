package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pangrams(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "check empty string",
			args: args{
				s: "",
			},
			want: "not pangram",
		},
		{
			name: "check for pangram",
			args: args{
				s: "We promptly judged antique ivory buckles for the next prize",
			},
			want: "pangram",
		},
		{
			name: "check for not pangram",
			args: args{
				s: "We promptly judged antique ivory buckles for the prize",
			},
			want: "not pangram",
		},
		{
			name: "check complicated cases",
			args: args{
				s: "qmExzBIJmdELxyOFWv LOCmefk TwPhargKSPEqSxzveiun",
			},
			want: "pangram",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pangrams(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
