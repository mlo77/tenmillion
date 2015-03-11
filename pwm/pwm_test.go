package pwm

import (
	"fmt"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	Pwm_set_loglevel(0)
	Pwm_setup(10, 0)
	Pwm_init_channel(0, 20000)
	Pwm_print_channel(0)

	Pwm_add_channel_pulse(0, 17, 0, 50)
	Pwm_add_channel_pulse(0, 17, 100, 50)
	Pwm_add_channel_pulse(0, 17, 200, 50)
	Pwm_add_channel_pulse(0, 17, 300, 50)

	time.Sleep(5 * time.Second)

	Pwm_clear_channel_gpio(0, 17)
	Pwm_shutdown()
	fmt.Println("ok")
}