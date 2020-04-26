package models

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"

	"github.com/markbates/pkger"
)

type Stop struct {
	ID        string `json:"id"`
	StationID string `json:"station_id"`
}

var stops = make(map[string]*Stop)

func init() {
	csvFile, _ := pkger.Open("/resources/stops.txt")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Skip first line (header)
	reader.Read()

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		stationID := line[9]

		// When a station ID is not provided, it's the same as the stop ID.
		if stationID == "" {
			stationID = line[0]
		}

		stops[line[0]] = &Stop{
			ID:        line[0],
			StationID: stationID,
		}
	}
}

func GetStop(id string) *Stop {
	return stops[id]
}
