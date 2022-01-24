package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
	"github.com/xuri/excelize/v2"
)

var db *mcb.DB

type master struct {
	Data []string `json:"data"`
	// Age        int      `json:"age"`
	// Profession string   `json:"profession"`
	// Hobbies    []string `json:"hobbies"`
	// Type       string   `json:"type"`
}

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
	var myData master

	form := make(url.Values, 0)
	form.Add("bucket", "royaltypool") //bucket and collection-> namespace:bucket.scope.collection
	form.Add("aid", "d009")           //document ID
	form.Add("data", tempRow[2])


	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

}
