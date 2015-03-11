package adapter 

import (
	"fmt"
	"../pwm"
)

var ref int = 0

func ServoListen(c chan float32, gpio int, dma int, done chan bool) {
	fmt.Println("adapter", ref)
    ref++
	initPwm(dma)
	for {
		n, stillOpened := <-c
        if stillOpened {
            n += 100
            n *= 5
            if n < 10 {
                n = 10
            }
            if n > 1000 {
                n = 1000
            } 
            fmt.Println(n)
            makePulse(dma, gpio, int(n))
        } else {
            closePwm(dma, gpio)
            done <- true
            return
        }
	}
}

func initPwm(dma int) {
	subcycleTime := 20000 // 10 ms
    pwm.Pwm_set_loglevel(0)
    if pwm.Pwm_is_setup() != 1 {
        pwm.Pwm_setup(10, 0)
    }
	pwm.Pwm_init_channel(dma, subcycleTime)
	pwm.Pwm_print_channel(dma)
}

func closePwm(dma, gpio int) {
    pwm.Pwm_clear_channel_gpio(dma, gpio)
    if ref == 0 {
        pwm.Pwm_shutdown()
    }
}

func makePulse(dma, gpio, width int) {
    pwm.Pwm_add_channel_pulse(dma, gpio, 0, width)   
}

    // def set_servo(self, gpio, pulse_width_us):
    //     """
    //     Sets a pulse-width on a gpio to repeat every subcycle
    //     (by default every 20ms).
    //     """
    //     # Make sure we can set the exact pulse_width_us
    //     _pulse_incr_us = _PWM.get_pulse_incr_us()
    //     if pulse_width_us % _pulse_incr_us:
    //         # No clean division possible
    //         raise AttributeError(("Pulse width increment granularity %sus "
    //                 "cannot divide a pulse-time of %sus") % (_pulse_incr_us,
    //                 pulse_width_us))

    //     # Initialize channel if not already done, else check subcycle time
    //     if _PWM.is_channel_initialized(self._dma_channel):
    //         _subcycle_us = _PWM.get_channel_subcycle_time_us(self._dma_channel)
    //         if _subcycle_us != self._subcycle_time_us:
    //             raise AttributeError(("Error: DMA channel %s is setup with a "
    //                     "subcycle_time of %sus (instead of %sus)") % \
    //                     (_subcycle_us, self._subcycle_time_us))
    //     else:
    //         init_channel(self._dma_channel, self._subcycle_time_us)

    //     # If this GPIO is already used, clear it first
    //     if self._gpios_used & 1 << gpio:
    //         clear_channel_gpio(self._dma_channel, gpio)
    //     self._gpios_used |= 1 << gpio

    //     # Add pulse for this GPIO
    //     add_channel_pulse(self._dma_channel, gpio, 0, \
    //             int(pulse_width_us / _pulse_incr_us))