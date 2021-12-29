// Go Api server
package validador

import (
	"testing"
)

// go test -v -run ^TestRegexEmail$
func TestRegexEmail(t *testing.T) {
	type args struct {
		nameregex string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_regex_email_1", args{"teste@teste.com"}, true},
		{"test_regex_email_2", args{"testeteste.com"}, false},
		{"test_regex_email_3", args{"teste121@teste.com"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexEmail(tt.args.nameregex); got != tt.want {
				t.Errorf("RegexEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -v -run ^TestRegexPhone$
func TestRegexPhone(t *testing.T) {
	type args struct {
		nameregex string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_regex_phone_1", args{"(21) 2621-4426"}, true},
		{"test_regex_phone_2", args{"(21) 2621-4426sda"}, false},
		{"test_regex_phone_3", args{"(21) 98677-8136"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexPhone(tt.args.nameregex); got != tt.want {
				t.Errorf("RegexPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
