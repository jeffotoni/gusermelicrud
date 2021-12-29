package validador

import "testing"

func TestIsCPF(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_is_cpf_", args{"890.917.297-59"}, true},
		{"test_is_cpf_", args{"440.286.363-53"}, true},
		{"test_is_cpf_", args{"440.286.463-53"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCPF(tt.args.input); got != tt.want {
				t.Errorf("IsCPF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"test_is_cnpj_", args{"49.981.405/0001-47"}, true},
		{"test_is_cnpj_", args{"65.791.674/0001-05"}, true},
		{"test_is_cnpj_", args{"65.791.674/0002-05"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCNPJ(tt.args.input); got != tt.want {
				t.Errorf("IsCNPJ() = %v, want %v", got, tt.want)
			}
		})
	}
}
