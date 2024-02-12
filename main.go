package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	pin := rpio.Pin(21)
	pin.Input()
	pin.PullUp()

	for {
		if pin.Read() == rpio.Low {
			fmt.Println("pin 1 low")
		}
		time.Sleep(time.Second)
	}

	// dev, err := spi.Open(&spi.Devfs{
	// 	Dev:      "/dev/spidev0.1",
	// 	Mode:     spi.Mode3,
	// 	MaxSpeed: 500000,
	// })
	// if err != nil {
	// 	panic(err)
	// }
	//     const KEY1_PIN: u8 = 21;
	// const KEY2_PIN: u8 = 20;
	// const KEY3_PIN: u8 = 16;

}
