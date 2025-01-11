package main

import (
	"ento-go/src/models"
	"image/png"
	"os"
)

func main() {
	goban := models.NewGoban7()
	goban.ChangeTheme(models.NewDarkGobanTheme())

	goban.PlaceBlack('A', 5)
	goban.PlaceWhite('A', 4)
	goban.PlaceBlack('A', 3)
	goban.PlaceWhite('D', 4)
	goban.PlaceBlack('B', 4)

	goban.Print()

	image := goban.GetImage()

	file, err := os.Create("tmp/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, *image)
	if err != nil {
		panic(err)
	}
}
