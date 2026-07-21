package main

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

func MToF(m Meters) Feet { return Feet(m * 3.2808) }

func FToM(f Feet) Meters { return Meters(f / 3.2808) }

func PToK(p Pounds) Kilograms { return Kilograms(p * 0.45359237) }

func KToP(k Kilograms) Pounds { return Pounds(k / 0.45359237) }
