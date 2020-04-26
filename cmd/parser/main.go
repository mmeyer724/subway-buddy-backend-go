package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

type Station struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Location   *Location `json:"location"`
	NorthLabel string    `json:"north_label"`
	SouthLabel string    `json:"south_label"`
}

type Stop struct {
	ID        string `json:"id"`
	StationID string `json:"station_id"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "stations",
			Usage: "parses stations.csv and returns json",
			Action: func(c *cli.Context) error {
				if !c.Args().Present() {
					fmt.Println("error: please provide path to stations.csv")
				}
				parseStations(c.Args().First())
				return nil
			},
		},
		{
			Name:  "stops",
			Usage: "parses stops.txt and returns json",
			Action: func(c *cli.Context) error {
				if !c.Args().Present() {
					fmt.Println("error: please provide path to stops.txt")
				}
				parseStops(c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func parseStations(path string) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Skip first line (header)
	reader.Read()

	var stations []Station
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

		stations = append(stations, Station{
			ID:   line[2],
			Name: line[5],
			Location: &Location{
				Latitude:  lat,
				Longitude: lng,
			},
			NorthLabel: line[11],
			SouthLabel: line[12],
		})
	}

	json, _ := json.Marshal(stations)
	fmt.Println(string(json))
}

func parseStops(path string) {
	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	// Skip first line (header)
	reader.Read()

	var stops []Stop
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

		stops = append(stops, Stop{
			ID:        line[0],
			StationID: stationID,
		})
	}

	json, _ := json.Marshal(stops)
	fmt.Println(string(json))
}
