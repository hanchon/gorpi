package main

import (
	"github.com/hanchon/gorpi/assets"
)

func main() {
	assets.PrintImg(assets.Player(), true)
	assets.PrintImg(assets.Dino(), false)
	run := assets.Run()
	run.Mirror = false
	assets.PrintSprite(run, 0)
	run.Mirror = true
	assets.PrintSprite(run, 0)
}
