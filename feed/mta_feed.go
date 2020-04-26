package feed

import (
	"buddy/models"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/MobilityData/gtfs-realtime-bindings/golang/gtfs"
	proto "github.com/golang/protobuf/proto"
)

var feedUrls = []string{
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-ace",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-bdfm",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-g",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-jz",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-nqrw",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-l",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs",
	"https://api-endpoint.mta.info/Dataservice/mtagtfsfeeds/nyct%2Fgtfs-7",
}

var client = &http.Client{}

func PullFeed() {
	var trainsByStation = make(map[string][]models.Train)

	for _, url := range feedUrls {
		retrieveData(url, trainsByStation)
	}

	for stationID, trains := range trainsByStation {
		if !models.StationExists(stationID) {
			continue
		}
		models.SetTrains(stationID, trains)
	}
}

func retrieveData(url string, trainsByStation map[string][]models.Train) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", os.Getenv("ACCESS_KEY"))

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	feed := gtfs.FeedMessage{}
	err = proto.Unmarshal(body, &feed)
	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range feed.Entity {
		tripUpdate := entity.TripUpdate
		if tripUpdate == nil {
			continue
		}

		trip := tripUpdate.Trip

		for _, update := range tripUpdate.StopTimeUpdate {
			stopID := update.StopId
			stop := models.GetStop(*stopID)
			if stop == nil {
				continue
			}

			var direction models.Direction
			if strings.HasSuffix(*stopID, "N") {
				direction = models.North
			} else if strings.HasSuffix(*stopID, "S") {
				direction = models.South
			} else {
				log.Panic("Unknown direction for stop: " + *stopID)
			}

			trainsByStation[stop.StationID] = append(
				trainsByStation[stop.StationID],
				models.Train{
					RouteID:     *trip.RouteId,
					Direction:   direction,
					ArrivalTime: update.Arrival.GetTime(),
				},
			)
		}
	}
}
