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
		{name: "1", args: args{s: "testK.txt"}, want: []string{"Go ddd", "Bravo abc", "Gopher aac", "Alpha fff", "Grin ssa", "Delta bca"}, wantErr: false},
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

func Test_sortK(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{

		{name: "2", args: args{s: "testK.txt", n: 2}, want: []string{"Gopher aac", "Bravo abc", "Delta bca", "Go ddd", "Alpha fff", "Grin ssa"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			got, err := sortK(open, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortK() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortK() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortN(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "3", args: args{s: "testN.txt"}, want: []string{"1,2,3,4,5,5", "1,2,4,6,8,9", "1,2,2,4,7,9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			got, err := sortN(open)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortR(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "4", args: args{s: "testR.txt"}, want: []string{"Grin", "Gopher", "Go", "Delta", "Bravo", "Alpha"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			if got, _ := sortR(open); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortU(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "5", args: args{s: "testU.txt"}, want: []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open, _ := openFile(tt.args.s)
			if got, _ := sortU(open); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortU() = %v, want %v", got, tt.want)
			}
		})
	}
}
