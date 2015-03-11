package pwm


// #cgo LDFLAGS: -L./ -lpwm
// #include "pwm.h"
import "C"
import "fmt"
import "time"

func Pwm_setup(pw_incr_us, hw int) int {
	n, err := C.setup(C.int(pw_incr_us), C.int(hw))
	if err != nil {
		fmt.Println("pb pwm setup", err)
	}
	return int(n)
}

func Pwm_shutdown() {
	C.shutdown()
}

func Pwm_set_loglevel(level int) {
	C.set_loglevel(C.int(level))
}

func Pwm_init_channel(channel, subcycle_time_us int) int {
	n, err := C.init_channel(C.int(channel), C.int(subcycle_time_us))
	if err != nil {
		fmt.Println("pb init_channel", err)
	}
	return int(n)
}

func Pwm_clear_channel(channel int) int {
	n, err := C.clear_channel(C.int(channel))
	if err != nil {
		fmt.Println("pb clear_channel", err)
	}
	return int(n)
}

func Pwm_clear_channel_gpio(channel, gpio int) int {
	n, err := C.clear_channel_gpio(C.int(channel), C.int(gpio))
	if err != nil {
		fmt.Println("pb clear_channel_gpio", err)
	}
	return int(n)
}

func Pwm_print_channel(channel int) int {
	n, err := C.print_channel(C.int(channel))
	if err != nil {
		fmt.Println("pb print_channel", err)
	}
	//fmt.Println(n)
	return int(n)
}

func Pwm_add_channel_pulse(channel, gpio, width_start, width int) int {
	n, err := C.add_channel_pulse(C.int(channel), C.int(gpio), C.int(width_start), C.int(width))
	if err != nil {
		fmt.Println("pb add_channel_pulse", err)
	}
	return int(n)
}

func Pwm_get_error_message() {
	msg, err := C.get_error_message()
	if err != nil {
		fmt.Println("pb get_error_message", err)
	}
	fmt.Println(msg)
}

func Pwm_set_softfatal(enabled int) {
	_, err := C.set_softfatal(C.int(enabled))
	if err != nil {
		fmt.Println("pb set_softfatal", err)
	}
}

func Pwm_is_setup() int {
	n, err := C.is_setup()
	if err != nil {
		fmt.Println("pb is_setup", err)
	}
	return int(n)
}

func Pwm_is_channel_initialized(channel int) int {
	n, err := C.is_channel_initialized(C.int(channel))
	if err != nil {
		fmt.Println("pb is_channel_initialized", err)
	}
	return int(n)
}

func Pwm_get_pulse_incr_us() int {
	n, err := C.get_pulse_incr_us()
	if err != nil {
		fmt.Println("pb get_pulse_incr_us", err)
	}
	return int(n)
}

func Pwm_get_channel_subcycle_time_us(channel int) int {
	n, err := C.get_channel_subcycle_time_us(C.int(channel))
	if err != nil {
		fmt.Println("pb get_channel_subcycle_time_us", err)
	}
	return int(n)
}

func Pwm_test() {
	Pwm_set_loglevel(0)
	Pwm_setup(10, 0)
	Pwm_init_channel(0, 20000)
	Pwm_print_channel(0)

	Pwm_add_channel_pulse(0, 17, 0, 50)
	
	time.Sleep(5 * time.Second)

	Pwm_clear_channel_gpio(0, 17)
	Pwm_shutdown()
	fmt.Println("ok")

}
