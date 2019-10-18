package userHandler

import "testing"

func Test_auth(t *testing.T) {
	type args struct {
		id  string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"auth", args{id: "super_admin", key: "123"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := auth(tt.args.id, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
