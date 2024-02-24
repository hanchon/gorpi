package main

import (
	"fmt"
	"time"

	"github.com/hanchon/gorpi/assets"
	"github.com/hanchon/gorpi/spi"
)

func main() {
	device := spi.NewDevice()
	defer device.CloseDevice()

	go func() {
		for {
			// TODO: now if you keep pressing the button it register more than 1 press
			// wait until the button is released before printing the next message
			if device.Button1.IsPressed() {
				fmt.Println("button 1 pressed")
			}
			if device.Button2.IsPressed() {
				fmt.Println("button 2 pressed")
			}
			if device.Button3.IsPressed() {
				fmt.Println("button 3 pressed")
			}
			if device.KeyUp.IsPressed() {
				fmt.Println("keyUp pressed")
			}
			if device.KeyDown.IsPressed() {
				fmt.Println("keyDown pressed")
			}
			if device.KeyLeft.IsPressed() {
				fmt.Println("keyLeft pressed")
			}
			if device.KeyRight.IsPressed() {
				fmt.Println("keyRight pressed")
			}
			if device.KeyPress.IsPressed() {
				fmt.Println("keyPress pressed")
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	sd := spi.NewScrenData()

	assets.ImgToScreenData(&assets.ConverterParams{
		Img:     assets.Player(),
		Reverse: true,
		Sd:      sd,
	})

	device.RenderScreen(sd)

	time.Sleep(10 * time.Second)
}
