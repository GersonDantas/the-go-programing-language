package main

import (
	"fmt"
	"os"
	"strconv"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Meters float64
type Feet float64
type Pounds float64
type Kilograms float64
type Weight float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	WinterK       Kelvin  = 300
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
func (m Meters) String() string     { return fmt.Sprintf("%gM", m) }
func (f Feet) String() string       { return fmt.Sprintf("%gF", f) }
func (p Pounds) String() string     { return fmt.Sprintf("%gP", p) }
func (k Kilograms) String() string  { return fmt.Sprintf("%gKg", k) }
func (w Weight) String() string     { return fmt.Sprintf("%gW", w) }

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"35"}
	}
	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cƒ: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		feet := Feet(t)
		meters := Meters(t)
		pounds := Pounds(t)
		kilograms := Kilograms(t)
		fmt.Println("Converting values for input:", t)
		fmt.Println("Kilograms:", kilograms)
		fmt.Println("Fahrenheit:", f)
		fmt.Println("Celsius:", c)
		fmt.Println("Feet:", feet)
		fmt.Println("Meters:", meters)
		fmt.Println("Pounds:", pounds)
		fmt.Println("Kilograms:", kilograms)
		fmt.Println("------------------------------")
	}
}

//!-
