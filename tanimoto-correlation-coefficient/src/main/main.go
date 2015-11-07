package main

import (
	"fmt"
	"image"
	"os"
	"com/steelzack/multicomparer"
)

func main() {
	fImg1, _ := os.Open("../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)

	fImg2, _ := os.Open("../mic-example-images/images2.jpeg")
	defer fImg2.Close()
	img2, _, _ := image.Decode(fImg2)

	fImg3, _ := os.Open("../mic-example-images/images3.jpeg")
	defer fImg3.Close()
	img3, _, _ := image.Decode(fImg3)

	fImg4, _ := os.Open("../mic-example-images/images4.jpeg")
	defer fImg4.Close()
	img4, _, _ := image.Decode(fImg4)


	ratio12 := getTCC(img1, img2, multicomparer.B)
	ratio34 := getTCC(img3, img4, multicomparer.B)

	fmt.Println("Ratio for 1 and 2:", ratio12)
	fmt.Println("Ratio for 3 and 4:", ratio34)
}

