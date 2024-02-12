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

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	button1 := initButton(key1Pin)
	button2 := initButton(key2Pin)
	button3 := initButton(key3Pin)

	keyUp := initButton(keyUpPin)
	keyDown := initButton(keyDownPin)
	keyLeft := initButton(keyLeftPin)
	keyRight := initButton(keyRightPin)
	keyPress := initButton(keyPressPin)

	for {
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
}
