// Copyright 2019 Alessandro Pomponio. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package conversion

import (
	"errors"
	"image/color"
	"testing"
)

func TestRGBAToHSV(t *testing.T) {

	type args struct {
		r uint32
		g uint32
		b uint32
		a uint32
	}
	tests := []struct {
		name  string
		args  args
		wantH float64
		wantS float64
		wantV float64
	}{
		{
			name:  "#bada55",
			wantH: 74,
			wantS: 61,
			wantV: 85,
		},
		{
			name:  "#7fe5f0",
			wantH: 186,
			wantS: 47,
			wantV: 94,
		},
		{
			name:  "#ff0000",
			wantH: 0,
			wantS: 100,
			wantV: 100,
		},
		{
			name:  "#ff80ed",
			wantH: 309,
			wantS: 50,
			wantV: 100,
		},
		{
			name:  "#696969",
			wantH: 0,
			wantS: 0,
			wantV: 41,
		},
		{
			name:  "#133337",
			wantH: 187,
			wantS: 65,
			wantV: 22,
		},
		{
			name:  "#065535",
			wantH: 156,
			wantS: 93,
			wantV: 33,
		},
		{
			name:  "#c0c0c0",
			wantH: 0,
			wantS: 0,
			wantV: 75,
		},
		{
			name:  "#5ac18e",
			wantH: 150,
			wantS: 53,
			wantV: 76,
		},
		{
			name:  "#666666",
			wantH: 0,
			wantS: 0,
			wantV: 40,
		},
		{
			name:  "#dcedc1",
			wantH: 83,
			wantS: 19,
			wantV: 93,
		},
		{
			name:  "#f7347a",
			wantH: 338,
			wantS: 79,
			wantV: 97,
		},
		{
			name:  "#000000",
			wantH: 0,
			wantS: 0,
			wantV: 0,
		},
		{
			name:  "#ffffff",
			wantH: 0,
			wantS: 0,
			wantV: 100,
		},
		{
			name:  "#ffc0cb",
			wantH: 350,
			wantS: 25,
			wantV: 100,
		},
	}

	for _, tt := range tests {

		testColor, err := ParseHexColorFast(tt.name)
		if err != nil {
			t.Errorf("RGBAToHSV() unable to parse color %s", tt.name)
		}

		t.Run(tt.name, func(t *testing.T) {

			gotH, gotS, gotV := RGBAToHSV(testColor.RGBA())

			if gotH != tt.wantH {
				t.Errorf("RGBAToHSV() Color %s: (%d %d %d %d)\ngotH = %v, want %v", tt.name, testColor.R, testColor.G, testColor.B, testColor.A, gotH, tt.wantH)
			}

			if gotS != tt.wantS {
				t.Errorf("RGBAToHSV() Color %s: (%d %d %d %d)\ngotS = %v, want %v", tt.name, testColor.R, testColor.G, testColor.B, testColor.A, gotS, tt.wantS)
			}

			if gotV != tt.wantV {
				t.Errorf("RGBAToHSV() Color %s: (%d %d %d %d)\ngotV = %v, want %v", tt.name, testColor.R, testColor.G, testColor.B, testColor.A, gotV, tt.wantV)
			}
		})
	}
}

func BenchmarkRGBAToHSV(b *testing.B) {

	// Hex:		#3a648c
	// RGBA:	14906 25700 35980 65535
	for i := 0; i < b.N; i++ {
		RGBAToHSV(14906, 25700, 35980, 65535)
	}

}

// from https://stackoverflow.com/a/54200713
var errInvalidFormat = errors.New("invalid format")

func ParseHexColorFast(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
