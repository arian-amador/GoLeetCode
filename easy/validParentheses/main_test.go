package main

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Empty s", args{""}, true},
		{"Small s", args{"["}, false},
		{"Small duplicate", args{"[["}, false},
		{"Valid ()", args{"()"}, true},
		{"Valid ()[]{}", args{"()[]{}"}, true},
		{"Valid {[]}", args{"{[]}"}, true},
		{"Valid [({(())}[()])]}", args{"[({(())}[()])]"}, true},
		{"Invalid ){", args{"){"}, false},
		{"Invalid ((", args{"(("}, false},
		{"Invalid {[])", args{"{[])"}, false},
		{"Invalid {[)}", args{"{[)}"}, false},
		{"Invalid ([)]", args{"([)]"}, false},
		{"Invalid [({]})]", args{"[({]})]"}, false},
		{"Random bytes", args{"[]{1}(3)[a]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_String(t *testing.T) {
	type fields struct {
		bytes []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Print", fields{[]byte{1, 2, 3}}, "length: 3 content: [1 2 3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				bytes: tt.fields.bytes,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("stack.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
