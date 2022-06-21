package route

import (
	"errors"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID string `json: "routeID"`
	CLientID string `json: "clientID"`
	Positions []Position `json: "position"`
}

type Position struct {
	Lat float64 `json: "lat"`
	Long float64 `json: "long"`
}

type PartialROutePosition struct {
	ID string `json:"routeID"`
	ClientID string`json:"clientID"`
	Position []float64 `json:"position"` 
	Finished bool `json: "finished"`
}

func (r *Route) LoadPositions() error  {
	if r.ID == "" {
		return errors.New(text: "route ud not informed")
	}
	f, err := os.Open( name: "destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), sep: ",") 
		lat, err, := strconv.ParseFloat(data[0], bitSize: 64)
		if err != nil {
			return nil, err
		}

		long, err, := strconv.ParseFloat(data[1], bitSize: 64)
		if err != nil {
			return nil, err
		}
		r.Positions = appned(r.Positions, Position {
			Lat: lat,
			Long: long,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() ([]string, error) { 
	var route PartialROutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.CLientID = r.CLientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = appned(result, string(jsonRoute))
	}
	return result, nil
}