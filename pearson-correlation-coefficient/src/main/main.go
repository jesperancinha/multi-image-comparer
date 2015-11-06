package main

import (
	// Needs to be here to avoid nil error: 
	// https://github.com/golang/go/issues/10389
	// http://www.goinggo.net/2014/09/go-compiler-nil-pointer-checks.html
	_ "image/jpeg"
	"fmt"
	"image"
	"math"
	"os"
)

const (
	R = iota
	G = iota
	B = iota
	A = iota
)

func main() {
	fImg1, _ := os.Open("images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)

	fImg2, _ := os.Open("images2.jpeg")
	defer fImg2.Close()
	img2, _, _ := image.Decode(fImg2)

	fImg3, _ := os.Open("images3.jpeg")
	defer fImg3.Close()
	img3, _, _ := image.Decode(fImg3)

	fImg4, _ := os.Open("images4.jpeg")
	defer fImg4.Close()
	img4, _, _ := image.Decode(fImg4)

	ratio12 := getPCC(img1, img2, B)
	ratio34 := getPCC(img3, img4, B)

	fmt.Println("Ratio for 1 and 2:", ratio12)
	fmt.Println("Ratio for 3 and 4:", ratio34)
}

func getPCC(img1 image.Image, img2 image.Image, channel int) float64 {
	average1 := getAveragePerChannel(img1, channel)
	average2 := getAveragePerChannel(img2, channel)

	ratio := getPCCDetail(img1, img2, average1, average2, channel)
	return ratio
}

// A Determinant
// B Denominator X
// C Denominator Y
func getPCCDetail(img1 image.Image, img2 image.Image, average1 float64, average2 float64, channel int) float64 {
	dotLength := img1.Bounds().Dx()
	dotHeight := img1.Bounds().Dy()
	dett := 0.0
	denX := 0.0
	denY := 0.0
	for x := 0; x < dotLength; x++ {
		for y := 0; y < dotHeight; y++ {
			diffX := getChannelValue(img1, x, y, channel) - average1
			diffY := getChannelValue(img2, x, y, channel) - average2
			dett += (diffX) * (diffY)
			denX += math.Pow(diffX, 2)
			denY += math.Pow(diffY, 2)
		}
	}
	return dett / (math.Sqrt(denX) * math.Sqrt(denY))
}

func getAveragePerChannel(img image.Image, channel int) float64 {
	dotLength := img.Bounds().Dx()
	dotHeight := img.Bounds().Dy()
	dotCount := dotLength * dotHeight
	dotAverage := 0.0

	for x := 0; x < dotLength; x++ {
		for y := 0; y < dotHeight; y++ {
			dotAverage += float64(getChannelValue(img, x, y, channel)) / float64(dotCount)
		}
	}
	return dotAverage
}

func getChannelValue(img image.Image, x int, y int, channel int) float64 {
	var channelValue uint32 = 0
	switch channel {
	case R:
		channelValue, _, _, _ = img.At(x, y).RGBA()

	case G:
		_, channelValue, _, _ = img.At(x, y).RGBA()

	case B:
		_, _, channelValue, _ = img.At(x, y).RGBA()

	case A:
		_, _, _, channelValue = img.At(x, y).RGBA()
	}
	return float64(channelValue)
}
