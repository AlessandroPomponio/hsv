// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package histogram

import (
	"image"
	"math"
	"runtime"

	"github.com/AlessandroPomponio/hsv/conversion"
)

// With64Bins returns a color histogram with 64 bins for the input image.
// The values in the bins will represent the percentage of pixels mapped
// to a certain Hue, Saturation and Value level. The output can be considered
// as two 32 bin histograms, where the second batch of 32 bins represents
// colors with Value above 50.
// It is VERY IMPORTANT TO NOTICE that the percentages are rounded, so the
// sum of all percentages may not be equal to 100.
// The Hue will be mapped to 8 levels, indexes {0,4,8,12,16,20,24,28}.
// The Saturation will be mapped to 4 levels, indexes H_level + {0,1,2,3}.
// The Value will be mapped to 2 levels, indexes H_level + S_level + {0,32}.
func With64Bins(img image.Image, roundType int) []float64 {

	bins := make([]float64, 64)
	xBound := img.Bounds().Dx()
	yBound := img.Bounds().Dy()

	for x := 0; x < xBound; x++ {

		for y := 0; y < yBound; y++ {

			h, s, v := conversion.RGBAToHSV(img.At(x, y).RGBA())

			// hueBin in [0,7].
			// Try to map hue in equally-sized
			// levels by dividing it for 360/7.
			hueBin := int(h / 51.42857142857143)

			// saturationBin in [0,3]
			// Try to map saturation in equally-sized
			// levels by dividing it for 100/3.
			saturationBin := int(s / 33.33333333333333)

			// valueBin in [0,1]
			// Try to map value in equally-sized
			// levels by dividing it for a value
			// that's just above 50.
			valueBin := int(v / 50.0000000001)
			//valueBin = 0

			index := 4*hueBin + saturationBin + 32*valueBin
			bins[index]++

		}

	}

	return normalize64BinsHistogram(roundType, xBound, yBound, bins)

}

// With64BinsConcurrent returns a color histogram with 64 bins for the input image.
// This concurrent version will check whether the image is taller or wider and iterate
// over the biggest dimension of the two, using use one goroutine per column/row.
// The values in the bins will represent the percentage of pixels mapped to a certain
// Hue, Saturation and Value level. The output can be considered  as two 32 bin histograms,
// where the second batch of 32 bins represents colors with Value above 50.
// It is VERY IMPORTANT TO NOTICE that the percentages are rounded, so the sum of all
// percentages may not be equal to 100.
// The Hue will be mapped to 8 levels, indexes {0,4,8,12,16,20,24,28}.
// The Saturation will be mapped to 4 levels, indexes H_level + {0,1,2,3}.
// The Value will be mapped to 2 levels, indexes H_level + S_level + {0,32}.
func With64BinsConcurrent(img image.Image, roundType int) []float64 {

	bins := make([]float64, 64)
	cpuAmt := runtime.NumCPU()
	binChannel := make(chan []float64, cpuAmt/2)

	// Split image into NumCPU sub-images to help speed up the computation.
	rectangles := splitInto(cpuAmt, img.Bounds())
	for _, rectangle := range rectangles {
		go calculate64BinsForRectangle(rectangle, img, binChannel)
	}

	// Gather the results from all goroutines and sum them.
	for i := 0; i < len(rectangles); i++ {

		currentBins := <-binChannel

		for i := 0; i < 64; i++ {
			bins[i] += currentBins[i]
		}

	}

	return normalize64BinsHistogram(roundType, img.Bounds().Dx(), img.Bounds().Dy(), bins)

}

func calculate64BinsForRectangle(rectangle image.Rectangle, img image.Image, outputChan chan []float64) {

	bins := make([]float64, 64)
	for x := rectangle.Min.X; x <= rectangle.Max.X; x++ {

		for y := rectangle.Min.Y; y <= rectangle.Max.Y; y++ {

			h, s, v := conversion.RGBAToHSV(img.At(x, y).RGBA())

			// hueBin in [0,7].
			// Try to map hue in equally-sized
			// levels by dividing it for 360/7.
			hueBin := int(h / 51.42857142857143)

			// saturationBin in [0,3]
			// Try to map saturation in equally-sized
			// levels by dividing it for 100/3.
			saturationBin := int(s / 33.333333333)

			// valueBin in [0,1]
			// Try to map value in equally-sized
			// levels by dividing it for a value
			// that's just above 50.
			valueBin := int(v / 50.0000000000001)

			index := 4*hueBin + saturationBin + 32*valueBin
			bins[index]++

		}
	}

	outputChan <- bins
}

// normalize64BinsHistogram normalizes 64-bin histograms by the amount of pixels in the image.
// It also makes sure that the sum of the percentages in bins[i] and bins[i+32] is equal to
// the rounded value of the percentage of their sum.
// bins[i] + bins[i+32] = round((bins[i] + bins[i+32]) * 100 / pixels)
func normalize64BinsHistogram(roundType int, width, height int, bins []float64) []float64 {

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

		totalPercentage := roundFunction((bins[i] + bins[i+32]) * 100 / pixels)
		firstHalfPercentage := roundFunction(bins[i] * 100 / pixels)
		bins[i] = firstHalfPercentage
		bins[i+32] = totalPercentage - firstHalfPercentage

	}

	return bins

}
