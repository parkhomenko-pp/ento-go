package models

import (
	"errors"
	"image/color"
)

const (
	light = 0
	dark  = 1

	woodLight = 2
	woodDark  = 3

	tgLight = 4
	tgDark  = 5
)

type GobanTheme struct {
	id uint8

	blackStoneFill   color.Color
	blackStoneStroke color.Color
	whiteStoneFill   color.Color
	whiteStoneStroke color.Color

	lastBlackStoneFill   color.Color
	lastBlackStoneStroke color.Color
	lastWhiteStoneFill   color.Color
	lastWhiteStoneStroke color.Color
}

func CreateGobanThemeById(id uint8) GobanTheme {
	switch id {
	case light:
		return NewLightGobanTheme()
	case dark:
		return NewDarkGobanTheme()
	case woodLight:
		return NewWoodLightGobanTheme()
	case woodDark:
		return NewWoodDarkGobanTheme()
	case tgLight:
		return NewTgLightGobanTheme()
	case tgDark:
		return NewTgDarkGobanTheme()
	}
	return NewLightGobanTheme()
}

func (t *GobanTheme) GetFilePathName() (string, error) {
	switch t.id {
	case light:
		return "light", nil
	case dark:
		return "dark", nil
	case woodLight:
		return "wood-light", nil
	case woodDark:
		return "wood-dark", nil
	case tgLight:
		return "tg-light", nil
	case tgDark:
		return "tg-dark", nil
	default:
		return "", errors.New("invalid theme")
	}
}

func NewLightGobanTheme() GobanTheme {
	return GobanTheme{
		id: light,

		blackStoneFill:   color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
	}
}

func NewDarkGobanTheme() GobanTheme {
	return GobanTheme{
		id: dark,

		blackStoneFill:   color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
	}
}

func NewWoodLightGobanTheme() GobanTheme {
	return GobanTheme{
		id: woodLight,

		blackStoneFill:   color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x2C, G: 0x2C, B: 0x33, A: 0xFF},
	}
}

func NewWoodDarkGobanTheme() GobanTheme {
	return GobanTheme{
		id: woodDark,

		blackStoneFill:   color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x20, G: 0x20, B: 0x24, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xF3, G: 0xF3, B: 0xF3, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xD1, G: 0xD1, B: 0xD6, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x4A, G: 0x30, B: 0x2C, A: 0xFF},
	}
}

func NewTgLightGobanTheme() GobanTheme {
	return GobanTheme{
		id: tgLight,

		blackStoneFill:   color.RGBA{R: 0x43, G: 0x61, B: 0x79, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0x43, G: 0x61, B: 0x79, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0x43, G: 0x61, B: 0x79, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x43, G: 0x61, B: 0x79, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xF4, G: 0xF4, B: 0xF4, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xF4, G: 0xF4, B: 0xF4, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x43, G: 0x61, B: 0x79, A: 0xFF},
	}
}

func NewTgDarkGobanTheme() GobanTheme {
	return GobanTheme{
		id: tgDark,

		blackStoneFill:   color.RGBA{R: 0x0B, G: 0x14, B: 0x1F, A: 0xFF},
		blackStoneStroke: color.RGBA{R: 0xB1, G: 0xC3, B: 0xD5, A: 0xFF},
		whiteStoneFill:   color.RGBA{R: 0xB1, G: 0xC3, B: 0xD5, A: 0xFF},
		whiteStoneStroke: color.RGBA{R: 0xB1, G: 0xC3, B: 0xD5, A: 0xFF},

		lastBlackStoneFill:   color.RGBA{R: 0x0B, G: 0x14, B: 0x1F, A: 0xFF},
		lastBlackStoneStroke: color.RGBA{R: 0xB1, G: 0xC3, B: 0xD5, A: 0xFF},
		lastWhiteStoneFill:   color.RGBA{R: 0xB1, G: 0xC3, B: 0xD5, A: 0xFF},
		lastWhiteStoneStroke: color.RGBA{R: 0x18, G: 0x21, B: 0x2D, A: 0xFF},
	}
}
