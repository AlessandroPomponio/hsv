package histogram

import (
	"image"
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		r image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want []image.Rectangle
	}{
		{
			name: "1000x1000",
			args: args{r: image.Rect(0, 0, 1000, 1000)},
			want: []image.Rectangle{image.Rect(0, 0, 499, 1000), image.Rect(500, 0, 1000, 1000)},
		},
		{
			name: "333x333",
			args: args{r: image.Rect(0, 0, 333, 333)},
			want: []image.Rectangle{image.Rect(0, 0, 165, 333), image.Rect(166, 0, 333, 333)},
		},
		{
			name: "500x500 no 0,0",
			args: args{r: image.Rect(1, 1, 500, 500)},
			want: []image.Rectangle{image.Rect(1, 1, 249, 500), image.Rect(250, 1, 500, 500)},
		},
		{
			name: "747x915 no 0,0",
			args: args{r: image.Rect(1, 1, 747, 915)},
			want: []image.Rectangle{image.Rect(1, 1, 373, 915), image.Rect(374, 1, 747, 915)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitHorizontally() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitInto(t *testing.T) {
	type args struct {
		amount    int
		rectangle image.Rectangle
	}
	tests := []struct {
		name string
		args args
		want []image.Rectangle
	}{
		{
			name: "4 rectangles from 1000x1000",
			args: args{amount: 4, rectangle: image.Rect(0, 0, 1000, 1000)},
			want: []image.Rectangle{
				image.Rect(0, 0, 248, 1000),
				image.Rect(249, 0, 499, 1000),
				image.Rect(500, 0, 749, 1000),
				image.Rect(750, 0, 1000, 1000),
			},
		},
		{
			name: "8 rectangles from 1000x1000",
			args: args{amount: 8, rectangle: image.Rect(0, 0, 1000, 1000)},
			want: []image.Rectangle{
				image.Rect(0, 0, 123, 1000),
				image.Rect(124, 0, 248, 1000),
				image.Rect(249, 0, 373, 1000),
				image.Rect(374, 0, 499, 1000),
				image.Rect(500, 0, 623, 1000),
				image.Rect(624, 0, 749, 1000),
				image.Rect(750, 0, 874, 1000),
				image.Rect(875, 0, 1000, 1000),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitInto(tt.args.amount, tt.args.rectangle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitInto() = %v, want %v", got, tt.want)
			}
		})
	}
}
