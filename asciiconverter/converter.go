package asciiconverter

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/blackjack/webcam"
)

const asciiChars = "@#w$9876543210?!abc;:+=-,._"

func ConvertWebcamToASCII(width int, fps int) {
	cam, err := webcam.Open("/dev/video0")
	if err != nil {
		fmt.Println("Can't open webcam:", err)
		return
	}
	defer cam.Close()
	formats := cam.GetSupportedFormats()
	var format webcam.PixelFormat
	for f, desc := range formats {
		if strings.Contains(desc, "MJPEG") {
			format = f
			break
		}
	}
	cam.SetImageFormat(format, 640, 480)
	cam.StartStreaming()
	defer cam.StopStreaming()

	fmt.Println("Starting webcam... you should see yourself in ASCII in 2 seconds!")
	time.Sleep(2 * time.Second)
	for {
		cam.WaitForFrame(5)
		frameData, _ := cam.ReadFrame()
		img, err := jpeg.Decode(bytes.NewReader(frameData))
		if err != nil {
			continue
		}

		ascii := imageToASCII(img, width)

		fmt.Print("\033[2J\033[H")
		fmt.Print(ascii)

		time.Sleep(time.Second / time.Duration(fps))
	}
}

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
	defer file.Close()
	_, format, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("not an image")
		return
	}
	file.Seek(0, 0)
	if format == "gif" {
		gifImg, err := gif.DecodeAll(file)
		fmt.Println(gifImg)
		if err != nil {
			fmt.Println("error on decoding gif")
			return
		}
		gifToASCII(gifImg, widthOfImage)
	} else {
		img, _, err := image.Decode(file)
		fmt.Println(img)
		if err != nil {
			fmt.Println("error on decoding an image")
			return
		}
		asciiImage := imageToASCII(img, widthOfImage)
		println(asciiImage)
	}

	// if err != nil {
	// 	println("ERROR deconding image:", err)
	// 	return
	// }
	// bounds := img.Bounds()
	// width, height := bounds.Max.X, bounds.Max.Y
	// fmt.Printf("Image size: %dx%d\n", width, height)
	// imageToASCII(img, widthOfImage)
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
			gray := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 255
			charIndex := int(gray / 256 * float64(len(asciiChars)))
			if charIndex >= len(asciiChars) {
				charIndex = len(asciiChars) - 1
			}

			result += string(asciiChars[charIndex])
		}
		result += "\n"
	}
	return result
}

func gifToASCII(gifImg *gif.GIF, width int) {
	for {
		for i, frame := range gifImg.Image {
			fmt.Print("\033[2J")
			fmt.Print("\033[H")

			ascii := imageToASCII(frame, width)
			fmt.Print(ascii)

			delay := gifImg.Delay[i] * 10
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}

}
