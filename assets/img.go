package assets

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/hanchon/gorpi/spi"
	"github.com/spakin/netpbm"
)

var (
	//go:embed player.pbm
	player []byte

	//go:embed dino.pbm
	dino []byte

	//go:embed run.pbm
	run []byte
)

type ConverterParams struct {
	Img     *netpbm.Image
	Reverse bool
	Sd      *spi.ScreenData
	OffsetX uint8
	OffsetY uint8
}

func ImgToScreenData(params *ConverterParams) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if params.Reverse {
			return !temp
		}
		return temp
	}

	bounds := (*params.Img).Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			c, _, _, _ := (*params.Img).At(x, y).RGBA()
			params.Sd.SetPixel(uint8(x)+params.OffsetX, uint8(y)+params.OffsetY, cmp(c))
		}
	}
}

func PrintImg(img *netpbm.Image, reverse bool) {
	cmp := func(a uint32) bool {
		temp := false
		if a == 0 {
			temp = true
		}
		if reverse {
			return !temp
		}
		return temp
	}

	bounds := (*img).Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			c, _, _, _ := (*img).At(x, y).RGBA()
			if cmp(c) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func bytesToImg(raw []byte) *netpbm.Image {
	bytes := bytes.NewReader(raw)

	img, err := netpbm.Decode(bytes, &netpbm.DecodeOptions{
		Target: netpbm.PBM,
		Exact:  false, // Can accept grayscale or B&W too
	})
	if err != nil {
		panic(err)
	}
	return &img
}

func Player() *netpbm.Image {
	return bytesToImg(player)
}

func Run() *Sprite {
	return &Sprite{
		SpriteSheet:     bytesToImg(run),
		SpriteWidth:     13,
		SpriteHeight:    16,
		SpriteSeparator: 3,
		SpriteAmount:    6,
		ColorsReverse:   true,
	}
}

func Dino() *netpbm.Image {
	return bytesToImg(dino)
}
