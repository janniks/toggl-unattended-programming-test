package api

import (
	"reflect"
	"testing"
)

func Test_cardSequence(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantA []int64
	}{
		{args: args{n: 5}, wantA: []int64{0, 1, 2, 3, 4}},
		{args: args{n: 10}, wantA: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := cardSequence(tt.args.n); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("cardSequence() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func Test_parseCardCodes(t *testing.T) {
	var empty []int64

	type args struct {
		codeQuery string
	}
	tests := []struct {
		name    string
		args    args
		wantIds []int64
		wantErr bool
	}{
		{args: args{codeQuery: "AC,AD,AH"}, wantIds: []int64{0, 13, 26}, wantErr: false},
		{args: args{codeQuery: "abc"}, wantIds: empty, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIds, err := parseCardCodes(tt.args.codeQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCardCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIds, tt.wantIds) {
				t.Errorf("parseCardCodes() gotIds = %v, want %v", gotIds, tt.wantIds)
			}
		})
	}
}
