package fuzz

import "testing"

func TestEqual(t *testing.T) {
	type args struct {
		a []byte
		b []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Slices have the same length",
			args: args{a: []byte{102, 97, 108, 99}, b: []byte{102, 97, 108, 99}},
			want: true,
		},
		{
			name: "Slices don’t have the  same length",
			args: args{a: []byte{102, 97, 99}, b: []byte{102, 97, 108, 99}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
