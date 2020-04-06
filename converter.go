package main

import (
	"fmt"
	"math"
)

//Converter is our struct type
type Converter struct{}

const (
	centimeterLength = 0.01
	footLength       = 0.3048
	millisecond      = 1000
	second           = 60 * millisecond
	minute           = 60 * second
	celsius          = 32
	fahrenheit       = 5 / 9
)

//Feet is used for our unit type
type Feet float64

//Centimeter is used for our unit type
type Centimeter float64

//CentimeterToFeet is our method signature
func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	var Result = Feet(c / 30.48)
	return Result
}

//FeetToCentimeter is our method signature
func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	var Result = Centimeter(f * 30.48)
	return Result
}

//Minute is used for our unit type
type Minute float64

//Second is used for our unit type
type Second float64

//Millisecond is used for our unit type
type Millisecond float64

//MinuteToSecond is therefore
func (cvr Converter) MinuteToSecond(m Minute) Second {
	var Result = Second(m * 60)
	return Result
}

//SecondToMillisecond is
func (cvr Converter) SecondToMillisecond(s Second) Millisecond {
	var Result = Millisecond(s * 1000)
	return Result
}

//Celsius is used for our unit type
type Celsius float64

//Fahrenheit is used for our unit type
type Fahrenheit float64

//CelsuisToFahrenheit is
func (cvr Converter) CelsuisToFahrenheit(c Celsius) Fahrenheit {
	var Result = Fahrenheit(c*9/5) + 32
	return Result
}

//FahrenheitToCelsius is
func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	var Result = Celsius(f-32) * 5 / 9
	return Result
}

//Radian is used for our unit type
type Radian float64

//Degree is used for our unit type
type Degree float64

//RadianToDegree is
func (cvr Converter) RadianToDegree(rad Radian) Degree {
	var Result = Degree(rad*180) / math.Pi //const "pi" was imported which is equivalent to  3.14
	return Result
}

//DegreeToRadian is
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	var Result = Radian(d/180) * math.Pi //const "pi" was imported which is equivalent to  3.14
	return Result
}

//Kilogram is used for our unit type
type Kilogram float64

//Pounds is used for our unit type
type Pounds float64

//KilogramToPounds is
func (cvr Converter) KilogramToPounds(kg Kilogram) Pounds {
	var Result = Pounds(kg * 2.205) //KilogramToPounds is equivalent to 2.205
	return Result
}

//PoundsToKilogram is
func (cvr Converter) PoundsToKilogram(lbs Pounds) Kilogram {
	var Result = Kilogram(lbs / 2.205) //KilogramToPounds is equivalent to 2.205
	return Result
}

func main() {
	cvr := Converter{}
	//calling and printing
	fmt.Printf("15.80cm to Feet is: ")
	fmt.Println(cvr.CentimeterToFeet(15.80))

	fmt.Printf("15.80Ft to Centimeter is: ")
	fmt.Println(cvr.FeetToCentimeter(15.80))

	fmt.Printf("15.80mins to Second is: ")
	fmt.Println(cvr.MinuteToSecond(15.80))

	fmt.Printf("15.80secs to Millisecond is: ")
	fmt.Println(cvr.SecondToMillisecond(15.80))

	fmt.Printf("15.80C to Fahrenheit is: ")
	fmt.Println(cvr.CelsuisToFahrenheit(15.80))

	fmt.Printf("15.80F to Celsius is: ")
	fmt.Println(cvr.FahrenheitToCelsius(15.80))

	fmt.Printf("15.80Rad to Degree is: ")
	fmt.Println(cvr.RadianToDegree(15.80))

	fmt.Printf("15.80 Degrees to Radian is: ")
	fmt.Println(cvr.DegreeToRadian(15.80))

	fmt.Printf("15.80kg to lbs is: ")
	fmt.Println(cvr.KilogramToPounds(15.80))

	fmt.Printf("15.80lbs to kg is: ")
	fmt.Println(cvr.KilogramToPounds(15.80))
}
