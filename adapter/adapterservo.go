package adapter 

import (
	"fmt"
	"github.com/mlo77/tenmillion/pwm"
)

func ServoListen(c chan, gpio int) {
	fmt.Println("adapter")
	init(gpio)
	for {
		n := <-c
		fmt.Println(n)
	}
}

func init(gpio int) {
	subcycleTime := 20000 // 10 ms
	pwm.Pwm_init_channel(gpio, subcycleTime)
	pwm.Pwm_print_channel(gpio)
}
