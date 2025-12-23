package asciiconverter

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

const asciiChars = "@%#*+=-:. "

func ConvertImageToASCII(imageFilePath string, width int) {

	imageFilePath = strings.TrimSpace(imageFilePath)
	if filepath.Ext(imageFilePath) == "" {
		fmt.Println("No File Extenshion Found")
		return
	}
	GetImageData(imageFilePath, width)

}

func GetImageData(imageFilePath string, widthOfImage int) {
	file, err := os.Open(imageFilePath)
	if err != nil {
		fmt.Println("ERROR Opening file:", err)
		return
	}
	fmt.Println(file)
	defer file.Close()
	img, _, err := image.Decode(file)

	if err != nil {
		println("ERROR deconding image:", err)
		return
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Image size: %dx%d\n", width, height)
	imageToASCII(img, widthOfImage)
}

func imageToASCII(img image.Image, width int) string {
	bounds := img.Bounds()
	aspectRatio := float64(bounds.Dy()) / float64(bounds.Dx())
	height := int(float64(width) * aspectRatio * 0.5)
	result := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			imgX := x * bounds.Dx() / width
			imgY := y * bounds.Dy() / height
			r, g, b, _ := img.At(imgX, imgY).RGBA()

			// Heavy Math For me :)
			gray := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 300
			charIndex := int(gray / 256 * float64(len(asciiChars)))
			if charIndex >= len(asciiChars) {
				charIndex = len(asciiChars) - 1
			}

			result += string(asciiChars[charIndex])
		}
		result += "\n"
	}
	println(result)
	return result
}
