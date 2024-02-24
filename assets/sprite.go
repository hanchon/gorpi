package assets

import (
	"fmt"

	"github.com/hanchon/gorpi/spi"
	"github.com/spakin/netpbm"
)

type Sprite struct {
	SpriteSheet     *netpbm.Image
	SpriteWidth     int
	SpriteHeight    int
	SpriteSeparator int
	SpriteAmount    int
	Reverse         bool
}

func SpriteToScreenData(s *Sprite, offsetX, offsetY, index uint8, sd *spi.ScreenData) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if s.Reverse {
			return !temp
		}
		return temp
	}

	for y := 0; y < s.SpriteHeight; y++ {
		for x := 0; x < s.SpriteWidth; x++ {
			tempX := (s.SpriteSeparator + s.SpriteWidth) * int(index)
			c, _, _, _ := (*s.SpriteSheet).At(tempX+x, y).RGBA()
			sd.SetPixel(uint8(x)+offsetX, uint8(y)+offsetY, cmp(c))
		}
	}
}

func PrintSprite(s *Sprite, offsetX, offsetY, index uint8) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if s.Reverse {
			return !temp
		}
		return temp
	}

	for y := 0; y < s.SpriteHeight; y++ {
		for x := 0; x < s.SpriteWidth; x++ {
			offsetX := (s.SpriteSeparator + s.SpriteWidth) * int(index)
			c, _, _, _ := (*s.SpriteSheet).At(offsetX+x, y).RGBA()
			if cmp(c) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
