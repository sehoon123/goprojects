package main

import "fmt"

type Room uint8
const (
	MasterRoom Room = 1 << iota
	BathRoom
	Kitchen
	LivingRoom
)

func SetLight(rooms, room Room) Room {
	return rooms | room
}

func ResetLight(rooms, room Room) Room {
	return rooms &^ room
}

func IsLightOn(rooms, room Room) bool {
	return rooms & room != 0
}

func TurnLights(rooms Room) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("Master room is on")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("Bath room is on")
	}
	if IsLightOn(rooms, Kitchen) {
		fmt.Println("Kitchen is on")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("Living room is on")
	}
}

func main() {
	var rooms Room
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, Kitchen)

	rooms = ResetLight(rooms, MasterRoom)
	TurnLights(rooms)
}