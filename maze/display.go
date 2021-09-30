package maze

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"os/exec"
)

func (m *Maze) Display() error {
	wallCellThickness := m.wallThickness + m.cellThickness

	imgWidth := (m.wallThickness+m.cellThickness)*m.width + 1*m.wallThickness
	imgHeight := (m.wallThickness+m.cellThickness)*m.height + 1*m.wallThickness

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: imgWidth,
			Y: imgHeight,
		},
	})

	black := color.RGBA{A: 0xff}
	white := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	green := color.RGBA{R: 0xa2, G: 0xff, B: 0x9d, A: 0xff}

	getColor := func(x, y int) color.RGBA {
		xMod, yMod := x%wallCellThickness, y%wallCellThickness

		isBorder := x < m.wallThickness || imgWidth-m.wallThickness <= x || y < m.wallThickness || imgHeight-m.wallThickness <= y
		isCorner := 0 <= xMod && xMod < m.wallThickness && 0 <= yMod && yMod < m.wallThickness

		column, row := int(math.Floor(float64(x)/float64(wallCellThickness))), int(math.Floor(float64(y)/float64(wallCellThickness)))
		c := m.cell(column, row)

		isWall := (xMod < m.wallThickness && c.isWall(Left)) ||
			(xMod > wallCellThickness && c.isWall(Right)) ||
			(yMod < m.wallThickness && c.isWall(Top)) ||
			(yMod > wallCellThickness && c.isWall(Bottom))

		if isBorder || isCorner || isWall {
			return black
		}

		if c.isPath {
			return green
		}

		return white
	}

	for y := 0; y < img.Rect.Max.Y; y++ {
		for x := 0; x < img.Rect.Max.X; x++ {
			img.Set(x, y, getColor(x, y))
		}
	}

	f, err := os.CreateTemp("", "*.png")
	if err != nil {
		return fmt.Errorf("cannot create temporary image file: %w", err)
	}

	errEncode := png.Encode(f, img)
	errClose := f.Close()

	if errEncode != nil {
		return fmt.Errorf("could not encode PNG file: %w", errEncode)
	}

	if errClose != nil {
		return fmt.Errorf("could not close PNG file: %w", errClose)
	}

	if err = exec.Command("open", f.Name()).Run(); err != nil {
		return fmt.Errorf("could not open final PNG file: %w", err)
	}

	return nil
}
