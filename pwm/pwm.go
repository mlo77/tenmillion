package pwm


// #cgo LDFLAGS: -lpwm
// include "pwm.h"
import "C"
import "fmt"

func Pwm_setup(pw_incr_us, hw int) int {
	n, err := C.setup(pw_incr_us, hw)
	if err != nil {
		fmt.Println("pb pwm setup", err)
	}
	return n
}

func Pwm_shutdown() {
	C.shutdown()
}

func Pwm_set_loglevel(level int) {
	C.set_loglevel(level)
}

func Pwm_init_channel(channel, subcycle_time_us int) int {
	n, err := C.init_channel(channel, subcycle_time_us)
	if err != nil {
		fmt.Println("pb init_channel", err)
	}
	return n
}

func Pwm_clear_channel(channel int) int {
	n, err := C.clear_channel(channel)
	if err != nil {
		fmt.Println("pb clear_channel", err)
	}
	return n
}

func Pwm_clear_channel_gpio(channel, gpio int) int {
	n, err := C.clear_channel_gpio(channel, gpio)
	if err != nil {
		fmt.Println("pb clear_channel_gpio", err)
	}
	return n
}

func Pwm_print_channel(channel int) int {
	n, err := C.print_channel(channel)
	if err != nil {
		fmt.Println("pb print_channel", err)
	}
	return n
}

func Pwm_add_channel_pulse(channel, gpio, width_start, width int) int {
	n, err := C.add_channel_pulse(channel, gpio, width_start, width)
	if err != nil {
		fmt.Println("pb add_channel_pulse", err)
	}
	return n
}

func Pwm_get_error_message() string {
	msg, err := C.get_error_message()
	if err != nil {
		fmt.Println("pb get_error_message", err)
	}
	return msg
}

func Pwm_set_softfatal(enabled int) {
	_, err := C.set_softfatal(enabled)
	if err != nil {
		fmt.Println("pb set_softfatal", err)
	}
}

func Pwm_is_setup() int {
	n, err := C.is_setup()
	if err != nil {
		fmt.Println("pb is_setup", err)
	}
	return n
}

func Pwm_is_channel_initialized(channel int) int {
	n, err := C.is_channel_initialized(channel)
	if err != nil {
		fmt.Println("pb is_channel_initialized", err)
	}
	return n
}

func Pwm_get_pulse_incr_us() int {
	n, err := C.get_pulse_incr_us()
	if err != nil {
		fmt.Println("pb get_pulse_incr_us", err)
	}
	return n
}

func Pwm_get_channel_subcycle_time_us(channel int) int {
	n, err := C.get_channel_subcycle_time_us(channel)
	if err != nil {
		fmt.Println("pb get_channel_subcycle_time_us", err)
	}
	return n
}

