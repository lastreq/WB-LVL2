package main

import "testing"

func Test_execCommands(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "1", args: args{input: "cd"}, wantErr: true},
		{name: "2", args: args{input: "cd .."}, wantErr: false},
		{name: "3", args: args{input: "ps"}, wantErr: false},
		{name: "4", args: args{input: "ps|pwd"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := execCommands(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("execCommands() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
