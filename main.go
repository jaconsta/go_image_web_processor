package main

import (
	// "bytes"
	// "encoding/base64"
	"fmt"
	"image"
	"image/png"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func loadFromFile(filePath string) (img image.Image, err error) {
	// Load File
	storedFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Could not load file")
	}
	defer storedFile.Close()

	// Validate Image type
	_, imageType, err := image.Decode(storedFile)
	if err != nil {
		return nil, fmt.Errorf("Could not load validate file type")
	}

	// Dunno if just using image. Decode would work.
	storedFile.Seek(0, 0)
	if imageType == "png" {
		img, err = png.Decode(storedFile)
		if err != nil {
			return nil, fmt.Errorf("Could not load (Decode) file to image")
		}
	}
	if imageType == "jpeg" {
		img, err = jpeg.Decode(storedFile)
		if err != nil {
			return nil, fmt.Errorf("Could not load (Decode) file to image")
		}
	}


	return  img, nil
}

func loadFromString(base64String string) (img image.Image, err error) {
	// // Encode Base64
	// var buff bytes.Buffer
	// png.Encode(&buff, base64String)
	// encodedImage := base64.StdEncoding.EncodeToString(buff.Bytes())

	// unbased, _ := base64.StdEncoding.DecodeString(encodedImage)
	// copyByte := bytes.NewReader(unbased)
	// copyImage, err := png.Decode(copyByte)
	// if err != nil {
	// 	log.Fatal("Error converting image String.")
	// }

	return nil, nil
}

func saveToFile(img image.Image, filename string) error {
	outputImage, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Could not create destination file.")
	}
	defer outputImage.Close()

	png.Encode(outputImage, img)
	return nil
}

func main() {
	encodedImage, err := loadFromFile("cuadro02.jpg")
	if err != nil {
		log.Fatal(err)
	}

	var maxWidth uint = 200
	var maxHeight uint = 0  // To preserve aspect ratio 
	thumbnail := resize.Resize(maxWidth, maxHeight, encodedImage, resize.Lanczos3)
	err = saveToFile(thumbnail, "three_copy.png")
	if err != nil {
		log.Fatal(err)
	}
}
