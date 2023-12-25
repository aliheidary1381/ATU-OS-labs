package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {
	inputImagePath := "C:\\Users\\alial\\Downloads\\ATU.png"
	img, err := loadImage(inputImagePath)
	if err != nil {
		log.Fatal(err)
	}

	kernelSize := 5
	sigma := 1.0
	kernel := generateGaussianKernel(kernelSize, sigma)

	// Blur the image 16 times
	for i := 0; i < 16; i++ {
		img = applyGaussianBlur(img, kernel)
	}

	// Save the blurred image
	outputImagePath := "output.png"
	err = saveImage(img, outputImagePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image blurred and saved successfully.")
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return err
	}

	return nil
}

func generateGaussianKernel(size int, sigma float64) *mat.Dense {
	var kernel []float64
	center := size / 2
	scale := 2.0 * sigma * sigma

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			x := float64(i - center)
			y := float64(j - center)
			value := math.Exp(-((x*x + y*y) / scale))
			kernel = append(kernel, value)
		}
	}

	sum := floats.Sum(kernel)
	for i := range kernel {
		kernel[i] /= sum
	}

	return mat.NewDense(size, size, kernel)
}

func applyGaussianBlur(img image.Image, kernel *mat.Dense) image.Image {
	bounds := img.Bounds()
	result := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, a := convolution(img, kernel, x, y)
			result.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}

	return result
}

func convolution(img image.Image, kernel *mat.Dense, x, y int) (uint8, uint8, uint8, uint8) {
	bounds := img.Bounds()
	centerX, centerY := kernel.Dims()
	centerX /= 2
	centerY /= 2

	var r, g, b, a float64

	for i := 0; i < centerX; i++ {
		for j := 0; j < centerY; j++ {
			ix := x + i - centerX/2
			iy := y + j - centerY/2

			ix = clamp(ix, bounds.Min.X, bounds.Max.X-1)
			iy = clamp(iy, bounds.Min.Y, bounds.Max.Y-1)

			pr, pg, pb, pa := img.At(ix, iy).RGBA()

			kv := kernel.At(i, j)

			r += float64(pr) * kv
			g += float64(pg) * kv
			b += float64(pb) * kv
			a += float64(pa) * kv
		}
	}

	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
