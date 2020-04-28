package main

import (
	"fmt"
	"log"
)

type VehicleError string

func (e VehicleError) Error() string {
	return string(e)
}

type Vehicle interface {
	Move(distance float64) interface{}
	TankUp(quantity float64) interface{}
}

type Auto struct {
	CarBrand        string
	FuelLevel       float64
	FuelConsumption float64
	MaxFuelLevel    float64
}

func (auto *Auto) Move(distance float64) interface{} {
	disctanceFuelConsumption := distance / auto.FuelConsumption
	if auto.FuelLevel < disctanceFuelConsumption {
		return VehicleError("Too small fuel!!!")
	}
	auto.FuelLevel -= disctanceFuelConsumption
	log.Printf(
		"Moving to a distance of %.2f km was successful. Fuel level: %.2f",
		distance, auto.FuelLevel,
	)
	return nil
}

func (auto *Auto) TankUp(quantity float64) interface{} {
	if auto.FuelLevel + quantity > auto.MaxFuelLevel {
		return VehicleError("Too big fuel!!!")
	}
	auto.FuelLevel += quantity
	log.Printf(
		"Add fuel in quantity %.2f littres was successful. Fuel level: %.2f",
		quantity, auto.FuelLevel,
	)
	return nil
}

func Sum(x int, y int) int {
    return x + y
}

func main() {
	autoMark := Auto{"audi", 50.0, 8.5, 75.0}
	fmt.Println(autoMark)
	var t Vehicle = &autoMark
	res := t.Move(100.0)
	if res != nil {
		log.Println(res)
	}
	fmt.Println(autoMark)
	res = t.TankUp(100.0)
	if res != nil {
		log.Println(res)
	}
	fmt.Println(autoMark)
}
