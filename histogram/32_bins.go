// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package histogram provides methods to compute approximate
// color histograms in the HSV color space.
package histogram

import (
	"github.com/AlessandroPomponio/hsv/conversion"
	"image"
	"math"
	"runtime"
)

const (
	// RoundClosest will round to the closest value using math.Round
	RoundClosest = iota

	// RoundUp will round to the closest bigger value using math.Ceil
	RoundUp

	// RoundDown will round to the closest lower value using math.Trunc
	RoundDown
)

// With32Bins returns a color histogram with 32 bins for the input image.
// The values in the bins will represent the percentage of pixels mapped
// to a certain Hue and Saturation level.
// It is VERY IMPORTANT TO NOTICE that the percentages are rounded, so the
// sum of all percentages may not be equal to 100.
// The Hue will be mapped to 8 levels, indexes {0,4,8,12,16,20,24,28}.
// The Saturation will be mapped to 4 levels, indexes hue_level + {0,1,2,3}.
// The Value channel is not taken into consideration, as to give invariance
// to light intensity.
func With32Bins(img image.Image, roundType int) []float64 {

	bins := make([]float64, 32)
	xBound := img.Bounds().Dx()
	yBound := img.Bounds().Dy()

	for x := 0; x < xBound; x++ {

		for y := 0; y < yBound; y++ {

			h, s, _ := conversion.RGBAToHSV(img.At(x, y).RGBA())

			// hueBin in [0,7].
			// Try to map hue in equally-sized
			// levels by dividing it for 360/7.
			hueBin := int(h / 51.42857142857143)

			// saturationBin in [0,3]
			// Try to map saturation in equally-sized
			// levels by dividing it for 100/3.
			saturationBin := int(s / 33.33333333333333)

			index := 4*hueBin + saturationBin
			bins[index]++

		}

	}

	return normalize32BinsHistogram(roundType, xBound, yBound, bins)

}

// With32BinsConcurrent returns a color histogram with 32 bins for the input image.
// This concurrent version will check whether the image is taller or wider and iterate
// over the biggest dimension of the two, using use one goroutine per column/row.
// The values in the bins will represent the percentage of pixels mapped to a
// certain Hue and Saturation level.
// It is VERY IMPORTANT TO NOTICE that the percentages are rounded, so the
// sum of all percentages may not be equal to 100.
// The Hue will be mapped to 8 levels, indexes {0,4,8,12,16,20,24,28}.
// The Saturation will be mapped to 4 levels, indexes hue_level + {0,1,2,3}.
// The Value channel is not taken into consideration, as to give invariance to
// light intensity.
func With32BinsConcurrent(img image.Image, roundType int) []float64 {

	bins := make([]float64, 32)
	cpuAmt := runtime.NumCPU()
	binChannel := make(chan []float64, cpuAmt/2)

	// Split image into NumCPU sub-images to help speed up the computation.
	rectangles := splitInto(cpuAmt, img.Bounds())
	for _, rectangle := range rectangles {
		go calculate32BinsForRectangle(rectangle, img, binChannel)
	}

	// Gather the results from all goroutines and sum them.
	for i := 0; i < len(rectangles); i++ {

		currentBins := <-binChannel

		for i := 0; i < 32; i++ {
			bins[i] += currentBins[i]
		}

	}

	return normalize32BinsHistogram(roundType, img.Bounds().Dx(), img.Bounds().Dy(), bins)

}

func calculate32BinsForRectangle(rectangle image.Rectangle, img image.Image, outputChan chan []float64) {

	bins := make([]float64, 32)
	for x := rectangle.Min.X; x <= rectangle.Max.X; x++ {

		for y := rectangle.Min.Y; y <= rectangle.Max.Y; y++ {

			h, s, _ := conversion.RGBAToHSV(img.At(x, y).RGBA())

			// hueBin in [0,7].
			// Try to map hue in equally-sized
			// levels by dividing it for 360/7.
			hueBin := int(h / 51.42857142857143)

			// saturationBin in [0,3]
			// Try to map saturation in equally-sized
			// levels by dividing it for 100/3.
			saturationBin := int(s / 33.33333333333333)

			index := 4*hueBin + saturationBin
			bins[index]++

		}

	}

	outputChan <- bins

}

// normalize32BinsHistogram normalizes 32-bin histograms by the amount of pixels in the image.
func normalize32BinsHistogram(roundType int, width, height int, bins []float64) []float64 {

	pixels := float64(width * height)
	var roundFunction func(x float64) float64

	switch roundType {
	case RoundClosest:
		roundFunction = math.Round
	case RoundUp:
		roundFunction = math.Ceil
	case RoundDown:
		roundFunction = math.Trunc
	default:
		return nil
	}

	for i := 0; i < 32; i++ {
		bins[i] = roundFunction(bins[i] * 100 / pixels)
	}

	return bins

}
