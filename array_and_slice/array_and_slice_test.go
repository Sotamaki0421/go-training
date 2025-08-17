package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLength3IntArray(t *testing.T) {
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
			got := GetLength3IntArray()
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

func TestAddElementToSlice(t *testing.T) {
	type args struct {
		input int
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "appendを使用してスライスに要素を加える",
			args: args{
				input: 3,
				slice: []int{1, 2},
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddElementToSlice(tt.args.input, tt.args.slice)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMakeSlice(t *testing.T) {
	type args struct {
		length   int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want []int
		appendFlag bool
	}{
		{
			name: "lengthを決めるとその分0埋めで返ってくる",
			args: args{
				length:   5,
				capacity: 5,
			},
			want: []int{0, 0, 0, 0, 0},
		},
		{
			name: "lengthをゼロにすると空で返ってくる",
			args: args{
				length:   0,
				capacity: 5,
			},
			want: []int{},
		},
		{
			name: "capacity以上に要素を詰める",
			args: args{
				length:   5,
				capacity: 5,
			},
			want: []int{0, 0, 0, 0, 0, 0},
			appendFlag: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeSlice(tt.args.length, tt.args.capacity)
			if tt.appendFlag {
				got = append(got, 0)	
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
