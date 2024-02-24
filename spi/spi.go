package spi

import (
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
const busClockSpeed = 8_000_000
const rstPin = 25
const dcPin = 24
const csPin = 8
const blPin = 18
const width = 132
const height = 64
const pageSize = 8

func initButton(pinID uint8) *rpio.Pin {
	pin := rpio.Pin(pinID)
	pin.Input()
	pin.PullUp()
	return &pin
}

func writeCommand(command ...byte) {
	pin := rpio.Pin(dcPin)
	pin.Low()
	rpio.SpiTransmit(command...)
}

func writeData(data ...byte) {
	pin := rpio.Pin(dcPin)
	pin.High()
	rpio.SpiTransmit(data...)
}

func reset() {
	pin := rpio.Pin(dcPin)
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
	rst := rpio.Pin(rstPin)
	rst.Output()
	dc := rpio.Pin(dcPin)
	dc.Output()
	cs := rpio.Pin(csPin)
	cs.Output()
	bl := rpio.Pin(blPin)
	bl.Output()

	cs.Low()
	bl.High()
	dc.Low()

	err := rpio.SpiBegin(rpio.Spi0)
	if err != nil {
		panic(err)
	}
	rpio.SpiSpeed(busClockSpeed)
	rpio.SpiChipSelect(0)
	rpio.SpiMode(0, 0)
}

func render() {
	dc := rpio.Pin(dcPin)
	for page := 0; page < pageSize; page++ {
		writeCommand(byte(0xB0 + page))
		writeCommand(byte(0x02))
		writeCommand(byte(0x10))
		time.Sleep(10 * time.Millisecond)
		dc.High()

		for index := 0; index < width; index++ {
			if index == width/2 && page == 0 {
				rpio.SpiTransmit(byte(0b11000000))
			} else if index == width/2 && page == 1 {
				rpio.SpiTransmit(byte(0b01000000))
			} else if index == width/2 && page == 2 {
				rpio.SpiTransmit(byte(0b00000000))
			} else if index == width/2 && page == 3 {
				rpio.SpiTransmit(byte(0b11100000))
			} else if index == width/2 && page == 4 {
				rpio.SpiTransmit(byte(0b11110000))
			} else if index == width/2 && page == 6 {
				rpio.SpiTransmit(byte(0b11111111))
			} else {
				rpio.SpiTransmit(byte(0x00))
			}
		}
	}
	dc.Low()
}

func renderScreen(sd *ScreenData) {
	dc := rpio.Pin(dcPin)
	matrix := sd.matrixToBytes()
	for page := 0; page < pageSize; page++ {
		writeCommand(byte(0xB0 + page))
		writeCommand(byte(0x02))
		writeCommand(byte(0x10))
		time.Sleep(10 * time.Millisecond)
		dc.High()

		for index := 0; index < width; index++ {
			rpio.SpiTransmit(matrix[index][page])
		}
	}
	dc.Low()
}

func setContrast() {
	writeCommand([]byte{130, 0x7F}...)
}
