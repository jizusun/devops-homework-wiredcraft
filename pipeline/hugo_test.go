package pipeline

import "testing"

func Test_checkArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"no argument", args{}, "", true},
		{"unknown argument", args{[]string{"foo"}}, "", true},
		{"known argument", args{[]string{"dev"}}, "dev", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkArgs(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
