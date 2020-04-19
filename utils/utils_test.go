package utils

import "testing"

type A struct {
}

func TestGetUniqueName(t *testing.T) {
	type args struct {
		instance interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test get unique name",
			args: args{instance: ""},
			want: "string",
		},
		{
			name: "test custom type",
			args: args{instance: &A{}},
			want: "*utils.A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFullUniqueName(tt.args.instance); got != tt.want {
				t.Errorf("GetFullUniqueName() = %v, want %v", got, tt.want)
			}
		})
	}
}
