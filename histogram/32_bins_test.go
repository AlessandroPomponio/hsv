// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package histogram

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

/***********************************************************************************************************************
 *																													   *
 *													UTILITIES														   *
 *																													   *
 ***********************************************************************************************************************
 */

func getImageByRelativePath(relativePath string) image.Image {

	path, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatalf("Error while trying to get absolute path for %s: %s", relativePath, err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error while trying to open file: %s", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("Unable to close file")
		}
	}()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatalf("Error while trying to decode file: %s", err)
	}

	return img

}

/***********************************************************************************************************************
 *																													   *
 *													TESTS															   *
 *																													   *
 ***********************************************************************************************************************
 */

func TestWith32Bins(t *testing.T) {

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
			want:      []float64{13, 30, 3, 0, 4, 13, 1, 0, 3, 5, 0, 0, 2, 2, 4, 1, 5, 0, 0, 0, 2, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 29, 3, 0, 6, 12, 2, 0, 3, 4, 1, 0, 2, 2, 2, 1, 5, 0, 0, 0, 2, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{35, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 6, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 20, 21, 9, 0, 3, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{28, 1, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 7, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0, 23, 20, 9, 0, 3, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 2, 6, 1, 0, 22, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 2, 6, 1, 0, 21, 54, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := With32Bins(tt.img, tt.roundType)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("With32Bins()\nGot: %v\nWanted: %v", got, tt.want)
			}

		})

	}
}

func TestWith32BinsConcurrent(t *testing.T) {

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
			want:      []float64{13, 30, 3, 0, 4, 13, 1, 0, 3, 5, 0, 0, 2, 2, 4, 1, 5, 0, 0, 0, 2, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 29, 3, 0, 6, 12, 2, 0, 3, 4, 1, 0, 2, 2, 2, 1, 5, 0, 0, 0, 2, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{35, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 6, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 20, 21, 9, 0, 3, 0, 0, 0},
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{28, 1, 0, 0, 2, 0, 0, 0, 2, 0, 0, 0, 7, 0, 0, 0, 2, 0, 0, 0, 1, 0, 0, 0, 23, 20, 9, 0, 3, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 2, 6, 1, 0, 22, 53, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
			want:      []float64{14, 2, 6, 1, 0, 21, 54, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			got := With32BinsConcurrent(tt.img, tt.roundType)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("With32BinsConcurrent() = %v, want %v", got, tt.want)
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

func BenchmarkWith32Bins(b *testing.B) {

	img := getImageByRelativePath(`../pictures/beach_medium.jpg`)

	for i := 0; i < b.N; i++ {
		With32Bins(img, RoundClosest)
	}

}

func BenchmarkWith32BinsConcurrent(b *testing.B) {

	img := getImageByRelativePath(`../pictures/beach_medium.jpg`)

	for i := 0; i < b.N; i++ {
		With32BinsConcurrent(img, RoundClosest)
	}

}
