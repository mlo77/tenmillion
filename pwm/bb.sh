#!/bin/bash

rm libpwm.a
rm pwm.o
gcc -c pwm.c -o pwm.o
ar rcs libpwm.a pwm.o