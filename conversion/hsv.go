// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conversion

import (
	"math"
)

// RGBAToHSV transforms a color in the RGBA color space into the HSV equivalent.
// The formulas used can be found on Wikipedia.
// https://en.wikipedia.org/wiki/HSL_and_HSV#Color_conversion_formulae
func RGBAToHSV(rValue, gValue, bValue, aValue uint32) (h, s, v float64) {

	// The RGBA color components are scaled by the Alpha value, as per:
	// https://golang.org/src/image/color/color.go?s=2394:2435#L21
	// Since we need RGB values in the [0-1] range, we need to divide
	// them by A, making sure it's not 0.
	if aValue == 0 {
		return h, s, v
	}

	a := float64(aValue)
	r := float64(rValue) / a
	g := float64(gValue) / a
	b := float64(bValue) / a

	maxValue := math.Max(r, math.Max(g, b))

	// They're all 0s
	if maxValue == 0 {
		return 0, 0, 0
	}

	minValue := math.Min(r, math.Min(g, b))
	delta := maxValue - minValue

	// Greyscale, only V can be != 0
	if delta == 0 {
		return 0, 0, math.Round(maxValue * 100)
	}

	//hue
	switch maxValue {
	case r:
		h = 60 * ((g - b) / delta)
	case g:
		h = 60 * (((b - r) / delta) + 2)
	case b:
		h = 60 * (((r - g) / delta) + 4)
	}

	if h < 0 {
		h += 360
	}

	h = math.Round(h)

	//saturation
	s = math.Round(100 * delta / maxValue)

	//value
	v = math.Round(maxValue * 100)
	return h, s, v

}
