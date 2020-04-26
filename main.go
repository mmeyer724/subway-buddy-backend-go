package main

import (
	"buddy/feed"
)

func main() {
	// stop := models.GetStop("R30S")
	// station := models.GetStation(stop.StationID)
	// fmt.Printf("%+v\n", station)

	// models.SetTrains(stop.StationID, []*models.Train{
	// 	&models.Train{
	// 		RouteID:     "R",
	// 		Direction:   models.South,
	// 		ArrivalTime: "haha",
	// 	},
	// 	&models.Train{
	// 		RouteID:     "Q",
	// 		Direction:   models.North,
	// 		ArrivalTime: "meater",
	// 	},
	// })

	// fmt.Printf("%+v\n", models.GetStation(stop.StationID))
	feed.PullFeed()
}
