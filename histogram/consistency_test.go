// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package histogram

import (
	"image"
	"testing"
)

func TestConsistencyForNonConcurrentAlgorithms(t *testing.T) {

	tests := []struct {
		name      string
		img       image.Image
		roundType int
	}{
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 4692x7030",
			img:       getImageByRelativePath(`../pictures/beach_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			with64BinsResult := With64Bins(tt.img, tt.roundType)
			with32BinsResult := With32Bins(tt.img, tt.roundType)

			for i := 0; i < 32; i++ {

				if with32BinsResult[i] != (with64BinsResult[i] + with64BinsResult[i+32]) {
					t.Errorf("ConsistencyForNonConcurrentAlgorithms()\n64 bin result: %v\n32 bin result: %v", with64BinsResult, with32BinsResult)
				}

			}

		})

	}

}

func TestConsistencyForConcurrentAlgorithms(t *testing.T) {

	tests := []struct {
		name      string
		img       image.Image
		roundType int
	}{
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 4692x7030",
			img:       getImageByRelativePath(`../pictures/beach_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Laura Stanley from Pexels
			name:      "Beach 1280x1917",
			img:       getImageByRelativePath(`../pictures/beach_medium.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 6000x3532",
			img:       getImageByRelativePath(`../pictures/tree_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Mohsin khan from Pexels
			name:      "Tree 1280x753",
			img:       getImageByRelativePath(`../pictures/tree_medium.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 2448x3264",
			img:       getImageByRelativePath(`../pictures/lobster_original.jpg`),
			roundType: RoundClosest,
		},
		{
			//Photo by Toa Heftiba Şinca from Pexels
			name:      "Lobster 1280x1706",
			img:       getImageByRelativePath(`../pictures/lobster_medium.jpg`),
			roundType: RoundClosest,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			with64BinsResult := With64BinsConcurrent(tt.img, tt.roundType)
			with32BinsResult := With32BinsConcurrent(tt.img, tt.roundType)

			for i := 0; i < 32; i++ {

				if with32BinsResult[i] != (with64BinsResult[i] + with64BinsResult[i+32]) {
					t.Errorf("ConsistencyForConcurrentAlgorithms()\n64 bin result: %v\n32 bin result: %v", with64BinsResult, with32BinsResult)
				}

			}

		})

	}

}
