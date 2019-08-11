// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package histogram

import (
	"image"
	"reflect"
	"testing"
)

/***********************************************************************************************************************
 *																													   *
 *													TESTS															   *
 *																													   *
 ***********************************************************************************************************************
 */

func TestWith64Bins(t *testing.T) {

	tests := []struct {
		name      string
		img       image.Image
		roundType int
		want      []float64
	}{
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 4692x7030",
			img:       getImageByRelativePath(`../pictures/beach_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{4, 7, 1, 0, 3, 12, 1, 0, 2, 5, 0, 0, 1, 2, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 23, 2, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 2, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{5, 6, 2, 0, 4, 11, 2, 0, 2, 4, 1, 0, 1, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 23, 1, 0, 2, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 2, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 14, 8, 0, 1, 0, 0, 0, 28, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 6, 7, 1, 0, 2, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{6, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 14, 8, 0, 1, 0, 0, 0, 22, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0, 9, 6, 1, 0, 2, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 2, 6, 1, 0, 22, 52, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 2, 6, 1, 0, 21, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := With64Bins(tt.img, tt.roundType)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("With64Bins()\nGot: %v\nWanted: %v", got, tt.want)
			}

		})

	}
}

func TestWith64BinsConcurrent(t *testing.T) {

	tests := []struct {
		name      string
		img       image.Image
		roundType int
		want      []float64
	}{
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 4692x7030",
			img:       getImageByRelativePath(`../pictures/beach_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{4, 7, 1, 0, 3, 12, 1, 0, 2, 5, 0, 0, 1, 2, 4, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 23, 2, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 2, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{5, 6, 2, 0, 4, 11, 2, 0, 2, 4, 1, 0, 1, 2, 2, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 23, 1, 0, 2, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 2, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 14, 8, 0, 1, 0, 0, 0, 28, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 5, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 6, 7, 1, 0, 2, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{6, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 14, 8, 0, 1, 0, 0, 0, 22, 0, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 6, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0, 9, 6, 1, 0, 2, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 2, 6, 1, 0, 22, 52, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 2, 6, 1, 0, 21, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := With64BinsConcurrent(tt.img, tt.roundType)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("With64BinsConcurrent() = %v, want %v", got, tt.want)
			}

		})

	}
}

/***********************************************************************************************************************
 *																													   *
 *												     BENCHMARKS														   *
 *																													   *
 ***********************************************************************************************************************
 */

func BenchmarkWith64Bins(b *testing.B) {

	img := getImageByRelativePath(`../pictures/beach_medium.jpg`)

	for i := 0; i < b.N; i++ {
		With64Bins(img, RoundClosest)
	}

}

func BenchmarkWith64BinsConcurrent(b *testing.B) {

	img := getImageByRelativePath(`../pictures/beach_medium.jpg`)

	for i := 0; i < b.N; i++ {
		With64BinsConcurrent(img, RoundClosest)
	}

}
