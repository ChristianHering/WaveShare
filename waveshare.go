package waveshare

import (
	"image"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	width  = 800
	height = 480

	resetPin = rpio.Pin(17)
	dcPin    = rpio.Pin(25)
	csPin    = rpio.Pin(8)
	busyPin  = rpio.Pin(24)
)

func init() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
}

// Exit releases SPI/GPIO control
func Exit() {
	rpio.SpiEnd(rpio.Spi0)
	rpio.Close()
}

// Initialize Wakes up and initializes an EPD for use
func Initialize() {
	resetPin.Mode(rpio.Output)
	dcPin.Mode(rpio.Output)
	csPin.Mode(rpio.Output)
	busyPin.Mode(rpio.Input)

	Reset()

	SendCommand(1) //POWER SETTING
	SendData(7)
	SendData(7)  //VGH=20V,VGL=-20V
	SendData(63) //VDH=15V
	SendData(63) //VDL=-15V

	SendCommand(4) //POWER ON
	time.Sleep(100 * time.Millisecond)
	WaitUntilIdle()

	SendCommand(0) //PANNEL SETTING
	SendData(31)   //KW-3f   KWR-2F      BWROTP 0f       BWOTP 1f

	SendCommand(97) //tres
	SendData(3)     //source 800
	SendData(32)
	SendData(1) //gate 480
	SendData(224)

	SendCommand(21)
	SendData(0)

	SendCommand(80) //VCOM AND DATA INTERVAL SETTING
	SendData(16)
	SendData(7)

	SendCommand(96) //TCON SETTING
	SendData(34)
}

// DisplayImage Loads and displays an image on an EPD
func DisplayImage(dimg image.Image) {
	SendCommand(19)
	SendImageData(dimg, width, height)
}

// DisplayPartialImage displays "dimg" at point the given point
// TODO func DisplayPartialImage(dimg image.Image, point image.Point) {}

// Sleep turns off your EPD to prevent display damage
func Sleep() {
	SendCommand(2) //power off
	WaitUntilIdle()
	SendCommand(7) //deep sleep
	SendData(165)
}

// SendImageData transmits the bits of an image
//
// If trying to render an image to the screen, use DisplayImage or DisplayPartialImage instead
func SendImageData(dimg image.Image, width int, height int) {
	var bitArray []int //This was sadly the best way I could think of to construct a byte with custom bits

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if len(bitArray) == 8 {
				b := (bitArray[0] * 128) + (bitArray[1] * 64) + (bitArray[2] * 32) + (bitArray[3] * 16) + (bitArray[4] * 8) + (bitArray[5] * 4) + (bitArray[6] * 2) + (bitArray[7] * 1)

				SendData(byte(b))

				bitArray = []int{}
			}

			r, _, _, _ := dimg.At(x, y).RGBA()

			if r < 128 {
				bitArray = append(bitArray, 1)
			} else {
				bitArray = append(bitArray, 0)
			}
		}
	}
	TurnOnDisplay()
}

// WaitUntilIdle Loops until the initialized EPD is availible
func WaitUntilIdle() {
	for {
		SendCommand(113)

		if busyPin.Read() == rpio.Low {
			time.Sleep(200 * time.Millisecond)

			break
		}

		time.Sleep(time.Millisecond) //Sleep briefly so we're not maxing a thread
	}
}

// TurnOnDisplay displays the image stored on our EPD
func TurnOnDisplay() {
	SendCommand(18)                    //DISPLAY REFRESH
	time.Sleep(100 * time.Millisecond) //!!!The delay here is necessary, 200uS at least!!!
	WaitUntilIdle()

}

// Reset Resets EPD (calling this will wake up the EPD)
func Reset() {
	resetPin.Write(rpio.High)
	time.Sleep(200 * time.Millisecond)
	resetPin.Write(rpio.Low)
	time.Sleep(2 * time.Millisecond)
	resetPin.Write(rpio.High)
	time.Sleep(200 * time.Millisecond)
}

// SendCommand Sends an arbitrary command to the display
func SendCommand(b byte) {
	dcPin.Write(rpio.Low)
	csPin.Write(rpio.Low)
	rpio.SpiTransmit(b)
	csPin.Write(rpio.High)
}

// SendData Sends an arbitrary byte of data to the display
func SendData(b byte) {
	dcPin.Write(rpio.High)
	csPin.Write(rpio.Low)
	rpio.SpiTransmit(b)
	csPin.Write(rpio.High)
}
