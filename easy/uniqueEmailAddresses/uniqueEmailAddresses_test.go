package uniqueEmailAddresses

import (
	"testing"
)

func Test_numUniqueEmails(t *testing.T) {
	type args struct {
		emails []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty", args{[]string{}}, 0},
		{"1 Valid", args{[]string{"test@test.com"}}, 1},
		{"2 Valid", args{[]string{"test@test.com", "test2@test.com"}}, 2},
		{"Matching", args{[]string{"test2@test.com", "test.2@test.com"}}, 1},
		{"Ignore + in local", args{[]string{"test2+test@test.com", "test.2+test@test.com"}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numUniqueEmails(tt.args.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
