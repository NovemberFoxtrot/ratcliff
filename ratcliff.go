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

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: <source> <target>")
	}

	sourcefile, err := os.Open(os.Args[1])
	checkerror(err)
	defer sourcefile.Close()

	targetfile, err := os.Open(os.Args[2])
	checkerror(err)
	defer targetfile.Close()

	sourceimage, _, err := image.Decode(sourcefile)
	targetimage, _, err := image.Decode(targetfile)

	sourcebounds := sourceimage.Bounds()
	targetbounds := targetimage.Bounds()

	if (sourcebounds.Min.Y != targetbounds.Min.Y) || (sourcebounds.Min.X != targetbounds.Min.X) || (sourcebounds.Max.Y != targetbounds.Max.Y) || (sourcebounds.Max.X != targetbounds.Max.X) {
		log.Fatalln("Images are not the same size pixel-wise!")
	}

	for y := sourcebounds.Min.Y; y < sourcebounds.Max.Y; y++ {
		for x := sourcebounds.Min.X; x < sourcebounds.Max.X; x++ {
			sr, sg, sb, sa := sourceimage.At(x, y).RGBA()
			tr, tg, tb, ta := targetimage.At(x, y).RGBA()

			if (sr != tr) || (sg != tg) || (sb != tb) || (sa != ta) {
				log.Fatalln("Ah! They are not the same!")
			}
		}
	}

	log.Println("Oh, I guess they are the same image.")
}
