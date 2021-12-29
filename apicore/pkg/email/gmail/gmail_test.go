// @autor @jeffotoni
// actived count
// https://myaccount.google.com/lesssecureapps
package gemail

import "testing"

func TestSendUser(t *testing.T) {
	type args struct {
		to         []string
		subject    string
		user_email string
		link_pass  string
		domain     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendUser(tt.args.to, tt.args.subject, tt.args.user_email, tt.args.link_pass, tt.args.domain); (err != nil) != tt.wantErr {
				t.Errorf("SendUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
