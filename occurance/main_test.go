package main

import (
	"reflect"
	"testing"
)

func Test_matchingStrings(t *testing.T) {
	type args struct {
		str     []string
		queries []string
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Testing string occurrence ",
			args: args{
				str:     []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
				queries: []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
			},
			want: []int32{1, 3, 4, 3, 2},
		},
		{
			name: "Testing string occurrence with empty query",
			args: args{
				str:     []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
				queries: []string{},
			},
			want: []int32{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchingStrings(tt.args.str, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, got, tt.want)
			}
		})
	}
}
