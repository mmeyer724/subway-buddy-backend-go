package main

import (
	"buddy/models"
	"fmt"
)

func main() {
	stop := models.GetStop("R30S")
	station := models.GetStation(stop.StationID)

	fmt.Printf("%+v\n", stop)
	fmt.Printf("%+v\n", station)
}
