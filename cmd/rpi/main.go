package main

import (
	"fmt"
	"time"

	"github.com/hanchon/gorpi/assets"
	"github.com/hanchon/gorpi/spi"
)

const Speed = 7

type Pos struct {
	X            int
	Y            int
	DeviceWidth  int
	DeviceHeight int
	Sprite       *assets.Sprite
}

func (p *Pos) MoveY(moveUp bool) {
	temp := p.Y + Speed
	if moveUp {
		temp = p.Y - Speed
	}
	if temp < 0 {
		p.Y = 0
		return
	}
	if temp+p.Sprite.SpriteHeight > p.DeviceHeight {
		p.Y = p.DeviceHeight - p.Sprite.SpriteHeight
		return
	}
	p.Y = temp
}

func (p *Pos) MoveX(moveRight bool) {
	temp := p.X - Speed
	if moveRight {
		temp = p.X + Speed
		p.Sprite.Mirror = false
	} else {
		p.Sprite.Mirror = true
	}
	if temp < 0 {
		p.X = 0
		return
	}
	if temp+p.Sprite.SpriteWidth+10 > p.DeviceWidth {
		p.X = p.DeviceWidth - p.Sprite.SpriteWidth - 10
		return
	}
	p.X = temp
}

func main() {
	device := spi.NewDevice()
	defer func() {
		if r := recover(); r != nil {
			device.CloseDevice()
		} else {
			device.CloseDevice()
		}
	}()

	run := true
	player := assets.Run()
	playerPos := Pos{
		X:            0,
		Y:            0,
		DeviceWidth:  device.Width,
		DeviceHeight: device.Height,
		Sprite:       player,
	}

	go func() {
		for {
			// TODO: now if you keep pressing the button it register more than 1 press
			// wait until the button is released before printing the next message
			if device.Button1.IsPressed() {
				fmt.Println("button 1 pressed")
				run = false
			}
			if device.Button2.IsPressed() {
				fmt.Println("button 2 pressed")
				run = false
			}
			if device.Button3.IsPressed() {
				fmt.Println("button 3 pressed")
				run = false
			}
			if device.KeyUp.IsPressed() {
				fmt.Println("keyUp pressed")
				playerPos.MoveY(true)
			}
			if device.KeyDown.IsPressed() {
				fmt.Println("keyDown pressed")
				playerPos.MoveY(false)

			}
			if device.KeyLeft.IsPressed() {
				fmt.Println("keyLeft pressed")
				playerPos.MoveX(false)
			}
			if device.KeyRight.IsPressed() {
				fmt.Println("keyRight pressed")
				playerPos.MoveX(true)
			}
			if device.KeyPress.IsPressed() {
				fmt.Println("keyPress pressed")
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	counter := uint8(0)
	go func() {
		for run {
			if counter+1 > uint8(player.SpriteAmount)-1 {
				counter = 0
			} else {
				counter++
			}
			time.Sleep(500 * time.Millisecond)
		}
	}()
	for run {
		sd := spi.NewScrenData()
		assets.SpriteToScreenData(player, uint8(playerPos.X), uint8(playerPos.Y), counter, sd)
		// assets.ImgToScreenData(&assets.ConverterParams{
		// 	Img:     player,
		// 	Reverse: true,
		// 	Sd:      sd,
		// 	OffsetX: uint8(playerPos.X),
		// 	OffsetY: uint8(playerPos.Y),
		// })
		device.RenderScreen(sd)
		time.Sleep(17 * time.Microsecond)
		counter++
	}

}
