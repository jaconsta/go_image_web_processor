package main

import (
	// "bytes"
	// "encoding/base64"
	"io/ioutil"
	"fmt"
	"image"
	"image/png"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

var sourceFolder = "./source_images"
var destinationFolder = "./processed_images"

func readFolderFiles(folderPath string) ([]string, error) {
	// Try Glob too. https://golang.org/src/path/filepath/match.go?s=5609:5664#L224
	files, err := ioutil.ReadDir(folderPath)
    if err != nil {
        return nil, err
    }

		var fileNames []string
    for _, f := range files {
			if !f.IsDir() && f.Name() != ".gitkeep"{
				fileNames = append(fileNames, f.Name())
			}
    }
		return fileNames, nil
}

func loadFromFile(fileName string) (img image.Image, err error) {
	fileLocation := filepath.Join(sourceFolder, fileName)
	// Load File
	storedFile, err := os.Open(fileLocation)
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
	}	else if imageType == "jpeg" {
		img, err = jpeg.Decode(storedFile)
		if err != nil {
			return nil, fmt.Errorf("Could not load (Decode) file to image")
		}
	}	else {
		return nil, nil
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
	fileLocation := filepath.Join(destinationFolder, filename)
	outputImage, err := os.Create(fileLocation)
	if err != nil {
		return fmt.Errorf("Could not create destination file.")
	}
	defer outputImage.Close()

	png.Encode(outputImage, img)
	return nil
}

func main() {
	imageList, err := readFolderFiles(sourceFolder)
	if err != nil {
		log.Fatal(err)
	}
	for _, fileName := range imageList {
		encodedImage, err := loadFromFile(fileName)
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
}
