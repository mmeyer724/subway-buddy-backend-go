package models

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"strconv"

	"github.com/markbates/pkger"
)

type Station struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Location   *Location `json:"location"`
	NorthLabel string    `json:"north_label"`
	SouthLabel string    `json:"south_label"`
	Trains     []Train
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

var stations = make(map[string]*Station)

func init() {
	csvFile, _ := pkger.Open("/resources/Stations.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Skip first line (header)
	reader.Read()

	// var stations []Station
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		lat, err := strconv.ParseFloat(line[9], 64)
		if err != nil {
			panic(err)
		}
		lng, err := strconv.ParseFloat(line[10], 64)
		if err != nil {
			panic(err)
		}

		stations[line[2]] = &Station{
			ID:   line[2],
			Name: line[5],
			Location: &Location{
				Latitude:  lat,
				Longitude: lng,
			},
			NorthLabel: line[11],
			SouthLabel: line[12],
		}
	}
}

func StationExists(id string) bool {
	_, exists := stations[id]
	return exists
}

func GetStation(id string) *Station {
	return stations[id]
}

func SetTrains(id string, trains []Train) {
	stations[id].Trains = trains
}
