package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func checkerror(err error) {
	if err != nil {
		log.Fatalln("[ERROR]", err)
	}
}

func pixels(source, target string) bool {
	sourcefile, err := os.Open(source)
	checkerror(err)
	defer sourcefile.Close()

	targetfile, err := os.Open(target)
	checkerror(err)
	defer targetfile.Close()

	sourceimage, _, err := image.Decode(sourcefile)
	targetimage, _, err := image.Decode(targetfile)

	sourcebounds := sourceimage.Bounds()
	targetbounds := targetimage.Bounds()

	if (sourcebounds.Min.Y != targetbounds.Min.Y) || (sourcebounds.Min.X != targetbounds.Min.X) || (sourcebounds.Max.Y != targetbounds.Max.Y) || (sourcebounds.Max.X != targetbounds.Max.X) {
		return false
		// log.Fatalln("Images are not the same size pixel-wise!")
	}

	for y := sourcebounds.Min.Y; y < sourcebounds.Max.Y; y++ {
		for x := sourcebounds.Min.X; x < sourcebounds.Max.X; x++ {
			sr, sg, sb, sa := sourceimage.At(x, y).RGBA()
			tr, tg, tb, ta := targetimage.At(x, y).RGBA()

			if (sr != tr) || (sg != tg) || (sb != tb) || (sa != ta) {
				return false
				// log.Fatalln("Ah! They are not the same!", x, y)
			}
		}
	}

	return true
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: filepath filepath filepath...")
	}

	for i := 1; i < len(os.Args)-1; i++ {
		for j := i + 1; j < len(os.Args); j++ {
			samePixels := pixels(os.Args[i], os.Args[j])
			if samePixels == true {
				log.Println(os.Args[i], os.Args[j])
			}
		}
	}
}
