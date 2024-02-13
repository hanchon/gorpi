package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const key1Pin = 21
const key2Pin = 20
const key3Pin = 16

const keyUpPin = 6
const keyDownPin = 19
const keyLeftPin = 5
const keyRightPin = 26
const keyPressPin = 13

func initButton(pinID uint8) *rpio.Pin {
	pin := rpio.Pin(pinID)
	pin.Input()
	pin.PullUp()
	return &pin
}

const BUS_CLOCK_SPEED = 8_000_000
const RST_PIN = 25
const DC_PIN = 24
const CS_PIN = 8
const BL_PIN = 18
const WIDTH = 132
const HEIGHT = 64

func writeCommand(command ...byte) {
	pin := rpio.Pin(DC_PIN)
	pin.Low()
	rpio.SpiTransmit(command...)
}

func writeData(data ...byte) {
	pin := rpio.Pin(DC_PIN)
	pin.High()
	rpio.SpiTransmit(data...)
}

func reset() {
	pin := rpio.Pin(DC_PIN)
	pin.High()
	time.Sleep(time.Millisecond * 100)
	pin.Low()
	time.Sleep(time.Millisecond * 100)
	pin.High()
	time.Sleep(time.Millisecond * 100)
}

func initDisplay() {
	reset()
	writeCommand(byte(0xAE))
	writeCommand(byte(0x02))
	writeCommand(byte(0x10))
	writeCommand(byte(0x40))
	writeCommand(byte(0x81))
	writeCommand(byte(0xA0))
	writeCommand(byte(0xC0))
	writeCommand(byte(0xA6))
	writeCommand(byte(0xA8))
	writeCommand(byte(0x3F))
	writeCommand(byte(0xD3))
	writeCommand(byte(0x00))
	writeCommand(byte(0xd5))
	writeCommand(byte(0x80))
	writeCommand(byte(0xD9))
	writeCommand(byte(0xF1))
	writeCommand(byte(0xDA))
	writeCommand(byte(0x12))
	writeCommand(byte(0xDB))
	writeCommand(byte(0x40))
	writeCommand(byte(0x20))
	writeCommand(byte(0x02))
	writeCommand(byte(0xA4))
	writeCommand(byte(0xA6))
	time.Sleep(100 * time.Millisecond)
	writeCommand(byte(0xAF)) // --turn on oled panel
}

func initSPI() {
	/* INITIALIZE GPIO */
	rst := rpio.Pin(RST_PIN)
	rst.Output()
	dc := rpio.Pin(DC_PIN)
	dc.Output()
	cs := rpio.Pin(CS_PIN)
	cs.Output()
	bl := rpio.Pin(BL_PIN)
	bl.Output()

	cs.Low()
	bl.High()
	dc.Low()

	err := rpio.SpiBegin(rpio.Spi0)
	if err != nil {
		panic(err)
	}
	rpio.SpiSpeed(BUS_CLOCK_SPEED)
	rpio.SpiChipSelect(0)
	rpio.SpiMode(0, 0)
}

func render() {
	for page := 0; page < 8; page++ {
		writeCommand(byte(0xB0 + page))
		writeCommand(byte(0x02))
		writeCommand(byte(0x10))
		time.Sleep(10 * time.Millisecond)
		dc := rpio.Pin(DC_PIN)
		dc.High()

		for index := 0; index < WIDTH; index++ {
			// let byte = self.memory[index + self.width as usize * page as usize];
			rpio.SpiTransmit(byte(0xFF))
		}
	}
}

func display() {
	rpio.SpiBegin(rpio.Spi0)

}

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}
	defer rpio.Close()

	button1 := initButton(key1Pin)
	button2 := initButton(key2Pin)
	button3 := initButton(key3Pin)

	keyUp := initButton(keyUpPin)
	keyDown := initButton(keyDownPin)
	keyLeft := initButton(keyLeftPin)
	keyRight := initButton(keyRightPin)
	keyPress := initButton(keyPressPin)

	go func() {
		for {
			// TODO: now if you keep pressing the button it register more than 1 press
			// wait until the button is released before printing the next message
			if button1.Read() == rpio.Low {
				fmt.Println("button 1 pressed")
			}
			if button2.Read() == rpio.Low {
				fmt.Println("button 2 pressed")
			}
			if button3.Read() == rpio.Low {
				fmt.Println("button 3 pressed")
			}
			if keyUp.Read() == rpio.Low {
				fmt.Println("keyUp pressed")
			}
			if keyDown.Read() == rpio.Low {
				fmt.Println("keyDown pressed")
			}
			if keyLeft.Read() == rpio.Low {
				fmt.Println("keyLeft pressed")
			}
			if keyRight.Read() == rpio.Low {
				fmt.Println("keyRight pressed")
			}
			if keyPress.Read() == rpio.Low {
				fmt.Println("keyPress pressed")
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	initSPI()
	initDisplay()
	render()

	time.Sleep(60 * time.Second)
	rpio.SpiEnd(rpio.Spi0)
	fmt.Println("ending here")
}
