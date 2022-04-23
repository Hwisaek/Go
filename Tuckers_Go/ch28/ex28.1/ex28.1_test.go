package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_square(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{x: 9},
			want: 81,
		},
		{
			name: "",
			args: args{x: 3},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := square(tt.args.x); got != tt.want {
				t.Errorf("square() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSquare1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(3) should be 9")
}
