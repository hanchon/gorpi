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
	ColorsReverse   bool
	Mirror          bool
}

func SpriteToScreenData(s *Sprite, offsetX, offsetY, index uint8, sd *spi.ScreenData) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if s.ColorsReverse {
			return !temp
		}
		return temp
	}

	x := 0
	if s.Mirror {
		x = s.SpriteWidth - 1
	}

	iterator := func() bool {
		if s.Mirror {
			return x >= 0
		}
		return x < s.SpriteWidth
	}

	for y := 0; y < s.SpriteHeight; y++ {
		for iterator() {
			tempX := (s.SpriteSeparator + s.SpriteWidth) * int(index)
			c, _, _, _ := (*s.SpriteSheet).At(tempX+x, y).RGBA()
			if s.Mirror {
				sd.SetPixel(uint8(s.SpriteWidth-x)+offsetX, uint8(y)+offsetY, cmp(c))
				x--
			} else {
				sd.SetPixel(uint8(x)+offsetX, uint8(y)+offsetY, cmp(c))
				x++
			}
		}
		if s.Mirror {
			x = s.SpriteWidth
		} else {
			x = 0
		}
	}
}

func PrintSprite(s *Sprite, index uint8) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if s.ColorsReverse {
			return !temp
		}
		return temp
	}

	x := 0
	if s.Mirror {
		x = s.SpriteWidth - 1
	}

	iterator := func() bool {
		if s.Mirror {
			return x >= 0
		}
		return x < s.SpriteWidth
	}

	for y := 0; y < s.SpriteHeight; y++ {
		for iterator() {
			tempX := (s.SpriteSeparator + s.SpriteWidth) * int(index)
			c, _, _, _ := (*s.SpriteSheet).At(tempX+x, y).RGBA()
			if cmp(c) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}

			if s.Mirror {
				x--
			} else {
				x++
			}
		}
		fmt.Println("")
		if s.Mirror {
			x = s.SpriteWidth
		} else {
			x = 0
		}
	}
}
