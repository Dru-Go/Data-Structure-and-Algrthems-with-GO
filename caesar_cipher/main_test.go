package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_caesarCipher(t *testing.T) {
	type args struct {
		s string
		k int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing with empty string",
			args: args{
				s: "",
				k: 3,
			},
			want: "",
		},
		{
			name: "testing with string",
			args: args{
				s: "bbbbbb",
				k: 2,
			},
			want: "dddddd",
		},
		{
			name: "testing with string",
			args: args{
				s: "middle-Outz",
				k: 2,
			},
			want: "okffng-Qwvb",
		},
		{
			name: "testing with complex string",
			args: args{
				s: "Always-Look-on-the-Bright-Side-of-Life",
				k: 5,
			},
			want: "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj",
		},
		{
			name: "testing with complex string",
			args: args{
				s: "abcdefghijklmnopqrstuvwxyz",
				k: 2,
			},
			want: "cdefghijklmnopqrstuvwxyzab",
		},
		{
			name: "testing with hello world string",
			args: args{
				s: "Hello_World!",
				k: 4,
			},
			want: "Lipps_Asvph!",
		},
		{
			name: "testing with large k",
			args: args{
				s: "159357lcfd",
				k: 98,
			},
			want: "159357fwzx",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := caesarCipher(tt.args.s, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_caesarDeCipher(t *testing.T) {
	type args struct {
		s string
		k int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing with string",
			args: args{
				s: "okffng-Qwvb",
				k: 2,
			},
			want: "middle-Outz",
		},
		{
			name: "testing with complicated string",
			args: args{
				s: "159357fwzx",
				k: 98,
			},
			want: "159357lcfd",
		},
		{
			name: "testing with complicated string",
			args: args{
				s: "Lipps_Asvph!",
				k: 4,
			},
			want: "Hello_World!",
		},
		{
			name: "testing with complex string",
			args: args{
				s: "cdefghijklmnopqrstuvwxyzab",
				k: 2,
			},
			want: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name: "testing with complex string",
			args: args{
				s: "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj",
				k: 5,
			},
			want: "Always-Look-on-the-Bright-Side-of-Life",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decipherCipher(tt.args.s, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
