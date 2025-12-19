package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter image path: ")
	imageFilePath, _ := reader.ReadString('\n')
	imageFilePath = strings.TrimSpace(imageFilePath)
	if filepath.Ext(imageFilePath) == "" {
		fmt.Println("No File Extenshion Found")
		return
	}

	file, err := os.Open(imageFilePath)
	if err != nil {
		fmt.Println("ERROR Opening file:", err)
		return
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		println("ERROR deconding image:", err)
		return
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	fmt.Printf("Image size: %dx%d\n", width, height)
}

func GetImageData(imageFilePath string) {

}
