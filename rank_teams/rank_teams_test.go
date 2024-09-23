package rank_teams

import (
	"testing"
)

func Test_rankTeams(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestCase 1",
			args: args{votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"}},
			want: "ACB",
		},
		{
			name: "TestCase 2",
			args: args{votes: []string{"WXYZ", "XYZW"}},
			want: "XWYZ",
		},
		{
			name: "TestCase 3",
			args: args{votes: []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}},
			want: "ZMNAGUEDSJYLBOPHRQICWFXTVK",
		},
		{
			name: "TestCase 3",
			args: args{votes: []string{"FVSHJIEMNGYPTQOURLWCZKAX", "AITFQORCEHPVJMXGKSLNZWUY",
				"OTERVXFZUMHNIYSCQAWGPKJL", "VMSERIJYLZNWCPQTOKFUHAXG", "VNHOZWKQCEFYPSGLAMXJIUTR",
				"ANPHQIJMXCWOSKTYGULFVERZ", "RFYUXJEWCKQOMGATHZVILNSP", "SCPYUMQJTVEXKRNLIOWGHAFZ",
				"VIKTSJCEYQGLOMPZWAHFXURN", "SVJICLXKHQZTFWNPYRGMEUAO", "JRCTHYKIGSXPOZLUQAVNEWFM",
				"NGMSWJITREHFZVQCUKXYAPOL", "WUXJOQKGNSYLHEZAFIPMRCVT", "PKYQIOLXFCRGHZNAMJVUTWES",
				"FERSGNMJVZXWAYLIKCPUQHTO", "HPLRIUQMTSGYJVAXWNOCZEKF", "JUVWPTEGCOFYSKXNRMHQALIZ",
				"MWPIAZCNSLEYRTHFKQXUOVGJ", "EZXLUNFVCMORSIWKTYHJAQPG", "HRQNLTKJFIEGMCSXAZPYOVUW",
				"LOHXVYGWRIJMCPSQENUAKTZF", "XKUTWPRGHOAQFLVYMJSNEIZC", "WTCRQMVKPHOSLGAXZUEFYNJI"}},
			want: "VWFHSJARNPEMOXLTUKICZGYQ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
