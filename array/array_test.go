package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLengeth3IntArray(t *testing.T) {
	tests := []struct {
		name string
		want [3]int
	}{
		{
			name: "正常系",
			want: [3]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLengeth3IntArray()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestGetIntSlice(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "スライスを取得",
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetIntSlice()
			assert.Equal(t, got, tt.want)
		})
	}
}
