package main

import (
	"go-telegram-bot/src/models"
	"image/png"
	"os"
)

func main() {
	goban := models.NewGoban7()
	//goban.ChangeTheme(models.NewDarkGobanTheme())

	goban.PlaceBlack('A', 3)
	goban.PlaceWhite('c', 4)

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
