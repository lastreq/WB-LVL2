package main

import (
	"reflect"
	"testing"
)

func Test_openFile(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "1", args: args{s: "test.txt"}, want: []string{"Go ddd", "Bravo abc", "Gopher aac", "Alpha fff", "Grin ssa", "Delta bca"}, wantErr: false},
		{name: "2", args: args{s: "file.txt"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := openFile(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_grepA(t *testing.T) {
	type args struct {
		s, regExpr string
		n          int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{name: "2", args: args{s: "test.txt", n: 2, regExpr: "Al"}, want: []string{"Grin ssa", "Delta bca"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			var sliceFlags [4]bool
			got := grepA(open, tt.args.n, tt.args.regExpr, sliceFlags)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grepB(t *testing.T) {
	type args struct {
		s, regExpr string
		n          int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{name: "2", args: args{s: "test.txt", n: 2, regExpr: "Al"}, want: []string{"Gopher aac", "Bravo abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			var sliceFlags [4]bool
			got := grepB(open, tt.args.n, tt.args.regExpr, sliceFlags)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_grepC(t *testing.T) {
	type args struct {
		s, regExpr string
		n          int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{name: "2", args: args{s: "test.txt", n: 2, regExpr: "Al"}, want: []string{"Grin ssa", "Delta bca", "Gopher aac", "Bravo abc"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			var sliceFlags [4]bool
			got := grepC(open, tt.args.n, tt.args.regExpr, sliceFlags)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grepC() got = %v, want %v", got, tt.want)
			}
		})
	}
}
