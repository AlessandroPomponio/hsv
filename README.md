# go-hsv

A library to convert RGBA colors to HSV and create approximate color histograms, written in Go.

## Installation
```
go get github.com/AlessandroPomponio/hsv
```

## Usage

``` Go
func main() {
	
	file, _ := os.Open("test.jpg")
	defer file.Close()

	img, _ := jpeg.Decode(file)

	fmt.Println("This library contains sequential and concurrent methods for creating histograms")
	histogramWith32BinsRC := histogram.With32Bins(img, histogram.RoundClosest)
	histogramWith32BinsRU := histogram.With32Bins(img, histogram.RoundUp)
	histogramWith32BinsRD := histogram.With32BinsConcurrent(img, histogram.RoundDown)

	fmt.Println("Histogram with 32 bins, rounded to the closest value:", histogramWith32BinsRC)
	fmt.Println("Histogram with 32 bins, rounded up:", histogramWith32BinsRU)
	fmt.Println("Histogram with 32 bins, rounded up:", histogramWith32BinsRD, "\n")

	fmt.Println("Want 64 bins? I've got you covered!")
	histogramWith64Bins := histogram.With64BinsConcurrent(img, histogram.RoundClosest)
	fmt.Println("Histogram with 64 bins, rounded to the closest value:", histogramWith64Bins, "\n")


	fmt.Println("It also allows you to convert colors from the RGBA color space to the HSV")
	h, s, v := conversion.RGBAToHSV(186, 218, 85, 255)
	fmt.Println("#bada55, RGBA: 186 218 85 255, HSV:", h,s,v)
	
}
```

## Benchmarks

Benchmarks can be found in the `histogram` package and are run on the `beach_medium.jpg` image (1280x1917).

```
BenchmarkWith32Bins-8             	       5	 244802280 ns/op
BenchmarkWith32BinsConcurrent-8   	      20	  53650080 ns/op
BenchmarkWith64Bins-8             	       5	 250596880 ns/op
BenchmarkWith64BinsConcurrent-8   	      20	  54100025 ns/op
```
