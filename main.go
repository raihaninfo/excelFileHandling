package main

import (
	"fmt"
	"os"

	"github.com/mateors/mcb"
	"github.com/xuri/excelize/v2"
)

var db *mcb.DB

// type master struct {
// 	Data []string `json:"data"`
// 	// Age        int      `json:"age"`
// 	// Profession string   `json:"profession"`
// 	// Hobbies    []string `json:"hobbies"`
// 	// Type       string   `json:"type"`
// }

func init() {

	db = mcb.Connect(host, username, password, false)

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {

	f, err := excelize.OpenFile("file1.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	var tempRow []string
	for _, row := range rows {
		tempRow = append(tempRow, row...)
	}
	fmt.Println(tempRow)

	//How to insert into couchbase bucket
	// var myData master

	d := db.Query(`INSERT INTO 'royaltypool'.raihan.client (KEY,VALUE)
VALUES ( "airline_4",
    { "callsign": "MY-AIR",
      "country": "United States",
      "iata": "Z1",
      "icao": "AQZ",
      "name": "80-My Air",
      "id": "4444",
      "type": "airline"} ),
VALUES ( "airline_4",
    { "callsign": "AIR-X",
      "country": "United States",
      "iata": "X1",
      "icao": "ARX",
      "name": "10-AirX",
      "id": "4445",
      "type": "airline"} )
RETURNING *;`)

	fmt.Println(d)

}
