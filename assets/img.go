package assets

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/spakin/netpbm"
)

var (
	//go:embed player.pbm
	player []byte

	//go:embed dino.pbm
	dino []byte
)

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

func Dino() *netpbm.Image {
	return bytesToImg(dino)
}
