// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package histogram

import (
	"image"
)

// splitInto splits a rectangle in up to amount parts.
func splitInto(amount int, rectangle image.Rectangle) []image.Rectangle {

	// Avoid infinite recursion in case amount ends up being odd
	if amount < 4 {
		return split(rectangle)
	}

	// Check if it's worth to stop the recursion
	xBound := rectangle.Dx()
	yBound := rectangle.Dy()
	if xBound < 400 && yBound < 400 {
		return split(rectangle)
	}

	rects := split(rectangle)
	return append(splitInto(amount/2, rects[0]), splitInto(amount/2, rects[1])...)

}

// split splits a Rectangle in two, horizontally.
func split(r image.Rectangle) []image.Rectangle {

	return []image.Rectangle{
		{

			Min: image.Point{
				X: r.Min.X,
				Y: r.Min.Y,
			},
			Max: image.Point{
				X: ((r.Max.X + r.Min.X) / 2) - 1,
				Y: r.Max.Y,
			},
		},
		{
			Min: image.Point{
				X: (r.Max.X + r.Min.X) / 2,
				Y: r.Min.Y,
			},
			Max: image.Point{
				X: r.Max.X,
				Y: r.Max.Y,
			},
		},
	}

}
