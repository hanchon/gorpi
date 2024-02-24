package spi

import (
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

type Device struct {
	Button1  Button
	Button2  Button
	Button3  Button
	KeyUp    Button
	KeyDown  Button
	KeyLeft  Button
	KeyRight Button
	KeyPress Button
	Height   int
	Width    int
}

func NewDevice() *Device {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	// Init buttons
	device := &Device{
		Button1:  Button{Button: initButton(key1Pin), Name: "1"},
		Button2:  Button{Button: initButton(key2Pin), Name: "2"},
		Button3:  Button{Button: initButton(key3Pin), Name: "3"},
		KeyUp:    Button{Button: initButton(keyUpPin), Name: "Up"},
		KeyDown:  Button{Button: initButton(keyDownPin), Name: "Down"},
		KeyLeft:  Button{Button: initButton(keyLeftPin), Name: "Left"},
		KeyRight: Button{Button: initButton(keyRightPin), Name: "Right"},
		KeyPress: Button{Button: initButton(keyPressPin), Name: "Press"},
		Height:   height,
		Width:    width,
	}
	initSPI()
	setContrast()
	initDisplay()
	time.Sleep(time.Second)
	time.Sleep(100 * time.Millisecond)

	return device
}

func (d *Device) CloseDevice() {
	time.Sleep(10 * time.Second)
	writeCommand(byte(0xAE)) // --turn off the screen
	rpio.SpiEnd(rpio.Spi0)
	rpio.Close()
}

func (d *Device) RenderScreen(sd *ScreenData) {
	renderScreen(sd)
}
