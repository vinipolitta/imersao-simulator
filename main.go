package main

import route2 "github.com/vinipolitta/FullCycle/tree/master/application/route"


func main()  {
	route := route2.Route {
		ID: "1",
		CLientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
}