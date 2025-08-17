package runeliteral

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerLiteral(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "正常系",
			want: 1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntegerLiteral()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestRuneLiteral(t *testing.T) {
	tests := []struct {
		name string
		want rune
	}{
		{
			name: "正常系",
			want: 'a',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RuneLiteral()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestFloatOperation(t *testing.T) {
	tests := []struct {
		name string
		want float32
	}{
		{
			name: "0.2 + 0.1 = 0.3ではない",
			want : 0.300000000000000004,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatOperation()
			assert.Equal(t, got, tt.want)
		})
	}
}
