package longestSubstringWithoutRepeatingCharacters

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{""}, 0},
		{" ", args{" "}, 1},
		{"bbbbb", args{"bbbbb"}, 1},
		{"abba", args{"abba"}, 2},
		{"au", args{"au"}, 2},
		{"dvdf", args{"dvdf"}, 3},
		{"abcabcbb", args{"abcabcbb"}, 3},
		{"pwwkew", args{"pwwkew"}, 3},
		{"abcdefabcdefghijklmabbb", args{"abcdefabcdefghijklmabbb"}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
