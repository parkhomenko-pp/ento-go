package models

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
	"unicode"
)

type Goban struct {
	theme GobanTheme

	size  uint8
	komi  float32
	count float32

	dots           [][]uint8
	lastStoneColor uint8
	lastI          uint8
	lastJ          uint8

	whiteCaptured uint16
	blackCaptured uint16

	dotsTerritory [][]uint8
}

const (
	empty = 0
	black = 1
	white = 2

	startSizePx     = 62
	rectangleSizePx = 146
	stoneRadPx      = 55
	lastStoneRadPx  = 18
)

func newGoban(size uint8, komi float32) *Goban {
	dots := make([][]uint8, size)
	for i := range dots {
		dots[i] = make([]uint8, size)
	}
	return &Goban{
		size:           size,
		dots:           dots,
		theme:          *NewLightGobanTheme(),
		lastStoneColor: 0,
		komi:           komi,
	}
}

func NewGobanBySize(size uint8) *Goban {
	switch size {
	case 7:
		return NewGoban7()
	case 9:
		return NewGoban9()
	case 11:
		return NewGoban11()
	case 13:
		return NewGoban13()
	case 19:
		return NewGoban19()
	}

	return nil
}

func NewGoban7() *Goban {
	return newGoban(7, 4.5)
}

func NewGoban9() *Goban {
	return newGoban(9, 5.5)
}

func NewGoban11() *Goban {
	return newGoban(11, 5.5)
}

func NewGoban13() *Goban {
	return newGoban(13, 6.5)
}

func NewGoban19() *Goban {
	return newGoban(19, 6.5)
}

func (g *Goban) SetDots(dots [][]uint8) {
	if dots == nil {
		return
	}

	g.dots = dots
}

func (g *Goban) SetLastColor(color uint8) {
	g.lastStoneColor = color
}

func (g *Goban) ChangeTheme(theme *GobanTheme) {
	g.theme = *theme
}

func (g *Goban) Print() {
	horizontalMarks := "  A B C D E F G H I J K L M N O P Q R S T"[0 : (g.size+1)*2]
	print(horizontalMarks)
	fmt.Printf("\tCount: %0.1f\n", g.count)
	for i, row := range g.dots {
		print(g.size-uint8(i), " ")
		for _, dot := range row {
			switch dot {
			case empty:
				print("· ")
			case black:
				print("⚫ ")
			case white:
				print("⚪ ")
			}
		}
		print(g.size - uint8(i))

		switch i {
		case 0:
			println("\tKomi: ", strconv.FormatFloat(float64(g.komi), 'f', 1, 32))
		case 2:
			println("\tBlack territory: ", g.CountBlack())
		case 3:
			println("\tWhite territory: ", g.CountWhite())
		case 5:
			println("\tWhite captured: ", g.whiteCaptured)
		case 6:
			println("\tBlack captured: ", g.blackCaptured)
		default:
			println()
		}
	}
	println(horizontalMarks)
}

func (g *Goban) place(j, i uint8, color uint8) {
	g.dots[j][i] = color
	g.lastI = j
	g.lastJ = i
	g.lastStoneColor = color

	g.removeStonesWithoutLiberties()
}

func (g *Goban) checkPoint(i, j, c uint8) error {
	if j >= g.size || i >= g.size {
		return errors.New("out of range")
	}
	if g.lastStoneColor == empty && c == white {
		return errors.New("first move must be black")
	}
	if j < 0 || j >= uint8(len(g.dots)) || i < 0 || i >= uint8(len(g.dots)) {
		return errors.New("out of range")
	}
	if g.dots[i][j] != empty {
		return errors.New("already placed")
	}
	if g.lastStoneColor == uint8(c) {
		return errors.New("cannot place same color twice")
	}

	// Temporarily place the stone
	g.dots[i][j] = c
	defer func() { g.dots[i][j] = empty }()

	// Check if the move captures any opponent's stones
	opponentColor := black
	if c == black {
		opponentColor = white
	}
	visited := make([][]bool, g.size)
	for i := range visited {
		visited[i] = make([]bool, g.size)
	}
	for _, neighbor := range [][2]uint8{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
		nx, ny := neighbor[0], neighbor[1]
		if nx < g.size && ny < g.size && g.dots[nx][ny] == uint8(opponentColor) {
			group, hasLiberties := g.findGroupAndLiberties(nx, ny, uint8(opponentColor), visited)
			if !hasLiberties {
				g.removeGroup(group)
			}
		}
	}

	// Check if the stone or group has at least one liberty
	visited = make([][]bool, g.size)
	for i := range visited {
		visited[i] = make([]bool, g.size)
	}
	_, hasLiberties := g.findGroupAndLiberties(i, j, c, visited)
	if !hasLiberties {
		return errors.New("move is suicidal")
	}

	return nil
}

func (g *Goban) isEmpty() bool {
	for _, row := range g.dots {
		for _, col := range row {
			if col != empty {
				return false
			}
		}
	}

	return true
}

func (g *Goban) letterToNumber(letter rune) (uint8, error) {
	if !unicode.IsLetter(letter) {
		return 0, errors.New("not a letter")
	}

	index := uint8(unicode.ToUpper(letter)) - 'A'

	if index >= g.size {
		return 0, errors.New("out of goban size")
	}

	return index, nil
}

func (g *Goban) PlaceBlack(s rune, i uint8) error {
	i--

	j, err := g.letterToNumber(s)
	if err != nil {
		return err
	}

	i = g.size - i - 1

	if err := g.checkPoint(i, j, black); err != nil {
		return err
	}
	g.place(i, j, black)
	return nil
}

func (g *Goban) PlaceWhite(s rune, i uint8) error {
	i--

	j, err := g.letterToNumber(s)
	if err != nil {
		return err
	}

	i = g.size - i - 1

	if err := g.checkPoint(i, j, white); err != nil {
		return err
	}
	g.place(i, j, white)
	return nil
}

func DrawCircle(img draw.Image, cx, cy, r int, col color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			dist := math.Sqrt(float64(x*x + y*y))
			if dist <= float64(r) {
				alpha := 1.0
				if dist > float64(r)-1 {
					alpha = float64(r) - dist
				}
				originalColor := img.At(cx+x, cy+y)
				r1, g1, b1, a1 := originalColor.RGBA()
				r2, g2, b2, a2 := col.RGBA()
				newR := uint8((float64(r1)*(1-alpha) + float64(r2)*alpha) / 256)
				newG := uint8((float64(g1)*(1-alpha) + float64(g2)*alpha) / 256)
				newB := uint8((float64(b1)*(1-alpha) + float64(b2)*alpha) / 256)
				newA := uint8((float64(a1)*(1-alpha) + float64(a2)*alpha) / 256)
				img.Set(cx+x, cy+y, color.RGBA{R: newR, G: newG, B: newB, A: newA})
			}
		}
	}
}

func (g *Goban) loadBackground() (image.Image, error) {
	filePathName, err := g.theme.GetFilePathName()

	sourceImageFile, err := os.Open(
		"media/gobans/" + strconv.Itoa(int(g.size)) + "-" + filePathName + ".png",
	)
	if err != nil {
		println(err.Error())
	}

	defer func(sourceImageFile *os.File) {
		err := sourceImageFile.Close()
		if err != nil {
			log.Fatal("cannot close goban background file")
		}
	}(sourceImageFile)

	return png.Decode(sourceImageFile)
}

func (g *Goban) GetImage() **image.RGBA {
	gobanImage, err := g.loadBackground()
	if err != nil {
		return nil
	}

	drawableImage := image.NewRGBA(gobanImage.Bounds())
	draw.Draw(drawableImage, gobanImage.Bounds(), gobanImage, image.Point{}, draw.Src)

	for i, row := range g.dots {
		for j, dot := range row {
			if dot == empty {
				continue
			}

			jPosition := startSizePx + (j)*rectangleSizePx
			iPosition := startSizePx + (i)*rectangleSizePx

			if dot == black {
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx,
					g.theme.blackStoneStroke,
				)
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx-2,
					g.theme.blackStoneFill,
				)
				if i == int(g.lastI) && j == int(g.lastJ) {
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx,
						g.theme.lastBlackStoneStroke,
					)
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx-2,
						g.theme.lastBlackStoneFill,
					)
				}
				continue
			}

			if dot == white {
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx,
					g.theme.whiteStoneStroke,
				)
				DrawCircle(
					drawableImage,
					jPosition, iPosition,
					stoneRadPx-2,
					g.theme.whiteStoneFill,
				)
				if i == int(g.lastI) && j == int(g.lastJ) {
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx,
						g.theme.lastWhiteStoneStroke,
					)
					DrawCircle(
						drawableImage,
						jPosition, iPosition,
						lastStoneRadPx-2,
						g.theme.lastWhiteStoneFill,
					)
				}
				continue
			}
		}
	}

	return &drawableImage
}

func (g *Goban) removeStonesWithoutLiberties() {
	visited := make([][]bool, g.size)
	for i := range visited {
		visited[i] = make([]bool, g.size)
	}

	for i := uint8(0); i < g.size; i++ {
		for j := uint8(0); j < g.size; j++ {
			if g.dots[i][j] != empty && !visited[i][j] {
				group, hasLiberties := g.findGroupAndLiberties(i, j, g.dots[i][j], visited)
				if !hasLiberties {
					g.removeGroup(group)
				}
			}
		}
	}
}

func (g *Goban) findGroupAndLiberties(i, j, color uint8, visited [][]bool) ([][2]uint8, bool) {
	var group [][2]uint8
	var stack [][2]uint8
	stack = append(stack, [2]uint8{i, j})
	hasLiberties := false

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := point[0], point[1]

		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		group = append(group, [2]uint8{x, y})

		neighbors := [][2]uint8{
			{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1},
		}

		for _, neighbor := range neighbors {
			nx, ny := neighbor[0], neighbor[1]
			if nx < g.size && ny < g.size {
				if g.dots[nx][ny] == empty {
					hasLiberties = true
				} else if g.dots[nx][ny] == color && !visited[nx][ny] {
					stack = append(stack, [2]uint8{nx, ny})
				}
			}
		}
	}

	return group, hasLiberties
}

func (g *Goban) removeGroup(group [][2]uint8) {
	for _, point := range group {
		x, y := point[0], point[1]
		if g.dots[x][y] == black {
			g.blackCaptured++
		} else if g.dots[x][y] == white {
			g.whiteCaptured++
		}
		g.dots[x][y] = empty
	}
}

func (g *Goban) countSurroundedPoints(_ uint8) int {
	return 0
}

func (g *Goban) CountBlack() int {
	return g.countSurroundedPoints(black)
}

func (g *Goban) CountWhite() int {
	return g.countSurroundedPoints(white)
}

func (g *Goban) GetTerritoriesCounts() (uint16, uint16) {
	// Initialize the territory array
	dots := make([][]uint8, g.size)
	for i := range dots {
		dots[i] = make([]uint8, g.size)
	}
	g.dotsTerritory = dots

	// Iterate over all points on the board
	for i := uint8(0); i < g.size; i++ {
		for j := uint8(0); j < g.size; j++ {
			if g.dots[i][j] == empty && g.dotsTerritory[i][j] == empty {
				// Check if the empty point is surrounded by one color
				group, c := g.findTerritoryGroup(i, j)
				if c != empty {
					for _, point := range group {
						g.dotsTerritory[point[0]][point[1]] = c
					}
				}
			}
		}
	}

	return g.getTerritoryCount(black), g.getTerritoryCount(white)
}

func (g *Goban) findTerritoryGroup(i, j uint8) ([][2]uint8, uint8) {
	var group [][2]uint8
	var stack [][2]uint8
	stack = append(stack, [2]uint8{i, j})
	visited := make([][]bool, g.size)
	for i := range visited {
		visited[i] = make([]bool, g.size)
	}
	c := empty
	isMixed := false

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		x, y := point[0], point[1]

		if visited[x][y] {
			continue
		}
		visited[x][y] = true
		group = append(group, [2]uint8{x, y})

		neighbors := [][2]uint8{
			{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1},
		}

		for _, neighbor := range neighbors {
			nx, ny := neighbor[0], neighbor[1]
			if nx < g.size && ny < g.size {
				if g.dots[nx][ny] == empty && !visited[nx][ny] {
					stack = append(stack, [2]uint8{nx, ny})
				} else if g.dots[nx][ny] != empty {
					if c == empty {
						c = int(g.dots[nx][ny])
					} else if c != int(g.dots[nx][ny]) {
						isMixed = true
					}
				}
			}
		}
	}

	if isMixed {
		c = empty
	}

	return group, uint8(c)
}

func (g *Goban) getTerritoryCount(color uint8) uint16 {
	count := uint16(0)

	for _, row := range g.dotsTerritory {
		for _, dot := range row {
			if dot == color {
				count++
			}
		}
	}

	return count
}

func (g *Goban) GetDots() [][]uint8 {
	return g.dots
}
