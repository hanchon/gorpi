package main

import (
	"github.com/hanchon/gorpi/assets"
)

func main() {
	assets.PrintImg(assets.Player(), true)
	assets.PrintImg(assets.Dino(), false)
}
