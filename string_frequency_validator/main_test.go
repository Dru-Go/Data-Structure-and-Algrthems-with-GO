package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testing empty string",
			args: args{
				s: "",
			},
			want: "YES",
		},
		{
			name: "Testing with valid string #1",
			args: args{
				s: "abc",
			},
			want: "YES",
		},
		{
			name: "Testing with valid string #2",
			args: args{
				s: "abcc",
			},
			want: "YES",
		},
		{
			name: "Testing with valid string #3",
			args: args{
				s: "abccc",
			},
			want: "NO",
		},
		{
			name: "Testing with valid string #4",
			args: args{
				s: "aabbcd",
			},
			want: "NO",
		},
		{
			name: "Testing with valid string #5",
			args: args{
				s: "abcdefghhgfedecba",
			},
			want: "YES",
		},
		{
			name: "Testing with valid string #6",
			args: args{
				s: "aabbc",
			},
			want: "YES",
		},
		{
			name: "Testing with valid string #7",
			args: args{
				s: "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd",
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
