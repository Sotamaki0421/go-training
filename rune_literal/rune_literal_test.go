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
