package main

import (
	"fmt"
	"time"

	"github.com/hanchon/gorpi/assets"
	"github.com/hanchon/gorpi/spi"
)

type Pos struct {
	X            int
	Y            int
	DeviceWidth  int
	DeviceHeight int
	PlayerWidth  int
	PlayerHeight int
}

func (p *Pos) MoveY(moveUp bool) {
	temp := p.Y - 1
	if moveUp {
		temp = p.Y + 1
	}
	if temp < 0 {
		p.Y = 0
		return
	}
	if temp+p.PlayerHeight > p.DeviceHeight {
		p.Y = p.PlayerHeight
		return
	}
	p.Y = temp
}

func (p *Pos) MoveX(moveRight bool) {
	temp := p.X - 1
	if moveRight {
		temp = p.Y + 1
	}
	if temp < 0 {
		p.X = 0
		return
	}
	if temp+p.PlayerWidth > p.DeviceWidth {
		p.Y = p.PlayerWidth
		return
	}
	p.Y = temp
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
	player := assets.Player()
	playerPos := Pos{
		X:            0,
		Y:            0,
		DeviceWidth:  device.Width,
		DeviceHeight: device.Height,
		PlayerWidth:  (*player).Bounds().Max.X,
		PlayerHeight: (*player).Bounds().Max.Y,
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

	for run {
		sd := spi.NewScrenData()
		assets.ImgToScreenData(&assets.ConverterParams{
			Img:     player,
			Reverse: true,
			Sd:      sd,
			OffsetX: uint8(playerPos.X),
			OffsetY: uint8(playerPos.Y),
		})
		device.RenderScreen(sd)
		time.Sleep(100 * time.Microsecond)
	}

}
