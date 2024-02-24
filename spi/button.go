package spi

import "github.com/stianeikeland/go-rpio/v4"

type Button struct {
	Name   string
	Button *rpio.Pin
}

func (b *Button) IsPressed() bool {
	return b.Button.Read() == rpio.Low
}
