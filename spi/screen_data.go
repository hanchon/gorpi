package spi

type ScreenData struct {
	data [width][height]bool
}

func NewScrenData() *ScreenData {
	return &ScreenData{}
}

func (sd *ScreenData) SetPixel(x uint8, y uint8, value bool) {
	sd.data[x][y] = value
}

func (sd *ScreenData) matrixToBytes() (res [width][pageSize]byte) {
	for kX, vX := range sd.data {
		for kY, vY := range vX {
			// TODO: make sure that this work for the last row
			row := kY / (pageSize - 1)
			if vY {
				res[kX][row] = res[kX][row] | byte(2<<(kY%8))
			}
		}
	}
	return res
}

// func (sd *ScreenData) addLetter(letter rune, initX uint8, initY uint8) {
// 	// TODO: validate that rune is in the array
// 	matrix, ok := text.LetterPixelArrays[letter]
// 	if !ok {
// 		fmt.Println("letter not found!")
// 		return
// 	}
// 	fmt.Println(matrix)
// 	for y, row := range matrix {
// 		for x, pixel := range row {
// 			if pixel {
// 				fmt.Println("1")
// 				// TODO: handle overflow
// 				sd.setPixel(initX+uint8(x), initY+uint8(y), true)
// 			} else {
// 				fmt.Println(" ")
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func (sd *ScreenData) addEmoji(initX uint8, initY uint8) {
// 	smily := []byte{0x7E, 0x81, 0x95, 0xB1, 0xB1, 0x95, 0x81, 0x7E}
// 	for y, v := range smily {
// 		for x := 0; x < 8; x++ {
// 			// Use bitwise AND operation to check each bit
// 			if (v & (1 << uint(7-x))) != 0 {
// 				sd.setPixel(initY+uint8(len(smily)-1-y), initX+uint8(x), true)
// 			}
// 		}
// 	}
//
// }
