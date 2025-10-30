package ui

import (
	"image"
	"image/png"
	"log"
	"os"
)

// LoadImage loads the images using the path
func LoadImage(path string) image.Image {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("ERROR: Failed to open image %s: %v\n", path, err)
	}
	img, err := png.Decode(f)
	if err != nil {
		log.Fatalf("ERROR: Failed to decode image %s: %v\n", path, err)
	}
	// Check if closing the file results in an error
	if err := f.Close(); err != nil {
		log.Fatalf("ERROR: Failed to close image file %s: %v\n", path, err)
	}
	return img
}

// GetCardImage grabs the path,loads images, adds to an array of images
func GetCardImage(currentHand []string) []image.Image {
	var images []image.Image

	//loops through the current hand calling each card, adding them as slices of images to the array
	for _, card := range currentHand {
		path := "deckImages/" + card + ".png"

		f, err := os.Open(path)
		if err != nil {
			log.Fatalf("ERROR: Failed to open image %s: %v\n", path, err)
		}

		img, err := png.Decode(f)
		if err != nil {
			log.Fatalf("ERROR: Failed to decode image %s: %v\n", path, err)
		}

		// Check if closing the file results in an error
		if err := f.Close(); err != nil {
			log.Fatalf("ERROR: Failed to close image file %s: %v\n", path, err)
		}

		images = append(images, img)
	}
	return images
}
