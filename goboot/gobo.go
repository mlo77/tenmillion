package goboot

import (
	_"time"
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

type PwmCmd struct {

}

func startBot(c chan int) {
	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor("raspi")
	led7 := gpio.NewLedDriver(r, "led", "7")
	led11 := gpio.NewLedDriver(r, "led", "11")
	led13 := gpio.NewLedDriver(r, "led", "13")
	led15 := gpio.NewLedDriver(r, "led", "15")

	// led7 := gpio.NewMotorDriver(r, "motor", "7")
	// led11 := gpio.NewMotorDriver(r, "motor", "11")
	// led13 := gpio.NewMotorDriver(r, "motor", "13")
	// led15 := gpio.NewMotorDriver(r, "motor", "15")

	//motor := gpio.NewMotorDriver(r, "motor", "3")

	work := func() {
		for {
			select {
				case d := <- c:
					fmt.Println("Toggling!")
					switch d {
					case 7:
						led7.On()
					case -7:
						led7.Off()	
					case 11:
						led11.On()
					case -11:
						led11.Off()
					case 13:
						led13.On()
					case -13:
						led13.Off()
					case 15:
						led15.On()
					case -15:
						led15.Off()
					}
			}
		}
	}

	// work := func() {
	// 	speed := byte(0)
	// 	fadeAmount := byte(15)

	// 	gobot.Every(200*time.Millisecond, func() {
	// 		led7.Speed(speed)
	// 		led11.Speed(speed)
	// 		led13.Speed(speed)
	// 		led15.Speed(speed)
	// 		speed = speed + fadeAmount
	// 		if speed == 0 || speed == 255 {
	// 			fadeAmount = -fadeAmount
	// 			}
	// 		})
	// }

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led7,led11,led13,led15},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
